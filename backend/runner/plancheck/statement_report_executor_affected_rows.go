package plancheck

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/pkg/errors"

	"github.com/bytebase/bytebase/backend/plugin/parser/sql/ast"
	storepb "github.com/bytebase/bytebase/proto/generated-go/store"
)

func getTableDataSize(metadata *storepb.DatabaseSchemaMetadata, schemaName, tableName string) int64 {
	if metadata == nil {
		return 0
	}
	for _, schema := range metadata.Schemas {
		if schema.Name != schemaName {
			continue
		}
		for _, table := range schema.Tables {
			if table.Name != tableName {
				continue
			}
			return table.RowCount
		}
	}
	return 0
}

func getAffectedRowsForPostgres(ctx context.Context, sqlDB *sql.DB, metadata *storepb.DatabaseSchemaMetadata, node ast.Node) (int64, error) {
	switch node := node.(type) {
	case *ast.InsertStmt, *ast.UpdateStmt, *ast.DeleteStmt:
		if node, ok := node.(*ast.InsertStmt); ok && len(node.ValueList) > 0 {
			return int64(len(node.ValueList)), nil
		}
		return getAffectedRowsCount(ctx, sqlDB, fmt.Sprintf("EXPLAIN %s", node.Text()), getAffectedRowsCountForPostgres)
	case *ast.AlterTableStmt:
		if node.Table.Type == ast.TableTypeBaseTable {
			schemaName := "public"
			if node.Table.Schema != "" {
				schemaName = node.Table.Schema
			}
			tableName := node.Table.Name

			return getTableDataSize(metadata, schemaName, tableName), nil
		}
		return 0, nil
	case *ast.DropTableStmt:
		var total int64
		for _, table := range node.TableList {
			schemaName := "public"
			if table.Schema != "" {
				schemaName = table.Schema
			}
			tableName := table.Name

			total += getTableDataSize(metadata, schemaName, tableName)
		}
		return total, nil

	default:
		return 0, nil
	}
}

func getAffectedRowsCountForPostgres(res []any) (int64, error) {
	// the res struct is []any{columnName, columnTable, rowDataList}
	if len(res) != 3 {
		return 0, errors.Errorf("expected 3 but got %d", len(res))
	}
	rowList, ok := res[2].([]any)
	if !ok {
		return 0, errors.Errorf("expected []any but got %t", res[2])
	}
	// test-bb=# EXPLAIN INSERT INTO t SELECT * FROM t;
	// QUERY PLAN
	// -------------------------------------------------------------
	//  Insert on t  (cost=0.00..1.03 rows=0 width=0)
	//    ->  Seq Scan on t t_1  (cost=0.00..1.03 rows=3 width=520)
	// (2 rows)
	if len(rowList) < 2 {
		return 0, errors.Errorf("not found any data")
	}
	// We need the row 2.
	rowTwo, ok := rowList[1].([]any)
	if !ok {
		return 0, errors.Errorf("expected []any but got %t", rowList[0])
	}
	// PostgreSQL EXPLAIN statement result has one column.
	if len(rowTwo) != 1 {
		return 0, errors.Errorf("expected one but got %d", len(rowTwo))
	}
	// Get the string value.
	text, ok := rowTwo[0].(string)
	if !ok {
		return 0, errors.Errorf("expected string but got %t", rowTwo[0])
	}

	rowsRegexp := regexp.MustCompile("rows=([0-9]+)")
	matches := rowsRegexp.FindStringSubmatch(text)
	if len(matches) != 2 {
		return 0, errors.Errorf("failed to find rows in %q", text)
	}
	value, err := strconv.ParseInt(matches[1], 10, 64)
	if err != nil {
		return 0, errors.Errorf("failed to get integer from %q", matches[1])
	}
	return value, nil
}

type affectedRowsCountExtractor func(res []any) (int64, error)

func getAffectedRowsCount(ctx context.Context, sqlDB *sql.DB, explainSQL string, extractor affectedRowsCountExtractor) (int64, error) {
	res, err := query(ctx, sqlDB, explainSQL)
	if err != nil {
		return 0, err
	}
	rowCount, err := extractor(res)
	if err != nil {
		return 0, errors.Wrapf(err, "failed to get affected rows count, res %+v", res)
	}
	return rowCount, nil
}

// Query runs the EXPLAIN or SELECT statements for advisors.
func query(ctx context.Context, connection *sql.DB, statement string) ([]any, error) {
	tx, err := connection.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	rows, err := tx.QueryContext(ctx, statement)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	columnNames, err := rows.Columns()
	if err != nil {
		return nil, err
	}

	columnTypes, err := rows.ColumnTypes()
	if err != nil {
		return nil, err
	}

	colCount := len(columnTypes)

	var columnTypeNames []string
	for _, v := range columnTypes {
		// DatabaseTypeName returns the database system name of the column type.
		// refer: https://pkg.go.dev/database/sql#ColumnType.DatabaseTypeName
		columnTypeNames = append(columnTypeNames, strings.ToUpper(v.DatabaseTypeName()))
	}

	data := []any{}
	for rows.Next() {
		scanArgs := make([]any, colCount)
		for i, v := range columnTypeNames {
			// TODO(steven need help): Consult a common list of data types from database driver documentation. e.g. MySQL,PostgreSQL.
			switch v {
			case "VARCHAR", "TEXT", "UUID", "TIMESTAMP":
				scanArgs[i] = new(sql.NullString)
			case "BOOL":
				scanArgs[i] = new(sql.NullBool)
			case "INT", "INTEGER":
				scanArgs[i] = new(sql.NullInt64)
			case "FLOAT":
				scanArgs[i] = new(sql.NullFloat64)
			default:
				scanArgs[i] = new(sql.NullString)
			}
		}

		if err := rows.Scan(scanArgs...); err != nil {
			return nil, err
		}

		rowData := []any{}
		for i := range columnTypes {
			if v, ok := (scanArgs[i]).(*sql.NullBool); ok && v.Valid {
				rowData = append(rowData, v.Bool)
				continue
			}
			if v, ok := (scanArgs[i]).(*sql.NullString); ok && v.Valid {
				rowData = append(rowData, v.String)
				continue
			}
			if v, ok := (scanArgs[i]).(*sql.NullInt64); ok && v.Valid {
				rowData = append(rowData, v.Int64)
				continue
			}
			if v, ok := (scanArgs[i]).(*sql.NullInt32); ok && v.Valid {
				rowData = append(rowData, v.Int32)
				continue
			}
			if v, ok := (scanArgs[i]).(*sql.NullFloat64); ok && v.Valid {
				rowData = append(rowData, v.Float64)
				continue
			}
			// If none of them match, set nil to its value.
			rowData = append(rowData, nil)
		}

		data = append(data, rowData)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return []any{columnNames, columnTypeNames, data}, nil
}

// OceanBaseQueryPlan represents the query plan of OceanBase.
type OceanBaseQueryPlan struct {
	ID       int    `json:"ID"`
	Operator string `json:"OPERATOR"`
	Name     string `json:"NAME"`
	EstRows  int64  `json:"EST.ROWS"`
	Cost     int    `json:"COST"`
	OutPut   any    `json:"output"`
}

// Unmarshal parses data and stores the result to current OceanBaseQueryPlan.
func (plan *OceanBaseQueryPlan) Unmarshal(data any) error {
	b, err := json.Marshal(data)
	if err != nil {
		return err
	}
	if b != nil {
		return json.Unmarshal(b, &plan)
	}
	return nil
}

func getAffectedRowsCountForOceanBase(res []any) (int64, error) {
	// the res struct is []any{columnName, columnTable, rowDataList}
	if len(res) != 3 {
		return 0, errors.Errorf("expected 3 but got %d", len(res))
	}
	rowList, ok := res[2].([]any)
	if !ok {
		return 0, errors.Errorf("expected []any but got %t", res[2])
	}
	if len(rowList) < 1 {
		return 0, errors.Errorf("not found any data")
	}

	plan, ok := rowList[0].([]any)
	if !ok {
		return 0, errors.Errorf("expected []any but got %t", rowList[0])
	}
	planString, ok := plan[0].(string)
	if !ok {
		return 0, errors.Errorf("expected string but got %t", plan[0])
	}
	var planValue map[string]json.RawMessage
	if err := json.Unmarshal([]byte(planString), &planValue); err != nil {
		return 0, errors.Wrapf(err, "failed to parse query plan from string: %+v", planString)
	}
	if len(planValue) > 0 {
		queryPlan := OceanBaseQueryPlan{}
		if err := queryPlan.Unmarshal(planValue); err != nil {
			return 0, errors.Wrapf(err, "failed to parse query plan from map: %+v", planValue)
		}
		if queryPlan.Operator != "" {
			return queryPlan.EstRows, nil
		}
		count := int64(-1)
		for k, v := range planValue {
			if !strings.HasPrefix(k, "CHILD_") {
				continue
			}
			child := OceanBaseQueryPlan{}
			if err := child.Unmarshal(v); err != nil {
				return 0, errors.Wrapf(err, "failed to parse field '%s', value: %+v", k, v)
			}
			if child.Operator != "" && child.EstRows > count {
				count = child.EstRows
			}
		}
		if count >= 0 {
			return count, nil
		}
	}
	return 0, errors.Errorf("failed to extract 'EST.ROWS' from query plan")
}

func getAffectedRowsCountForMysql(res []any) (int64, error) {
	// the res struct is []any{columnName, columnTable, rowDataList}
	if len(res) != 3 {
		return 0, errors.Errorf("expected 3 but got %d", len(res))
	}
	rowList, ok := res[2].([]any)
	if !ok {
		return 0, errors.Errorf("expected []any but got %t", res[2])
	}
	if len(rowList) < 1 {
		return 0, errors.Errorf("not found any data")
	}

	// MySQL EXPLAIN statement result has 12 columns.
	// the column 9 is the data 'rows'.
	// the first not-NULL value of column 9 is the affected rows count.
	//
	// mysql> explain delete from td;
	// +----+-------------+-------+------------+------+---------------+------+---------+------+------+----------+-------+
	// | id | select_type | table | partitions | type | possible_keys | key  | key_len | ref  | rows | filtered | Extra |
	// +----+-------------+-------+------------+------+---------------+------+---------+------+------+----------+-------+
	// |  1 | DELETE      | td    | NULL       | ALL  | NULL          | NULL | NULL    | NULL |    1 |   100.00 | NULL  |
	// +----+-------------+-------+------------+------+---------------+------+---------+------+------+----------+-------+
	//
	// mysql> explain insert into td select * from td;
	// +----+-------------+-------+------------+------+---------------+------+---------+------+------+----------+-----------------+
	// | id | select_type | table | partitions | type | possible_keys | key  | key_len | ref  | rows | filtered | Extra           |
	// +----+-------------+-------+------------+------+---------------+------+---------+------+------+----------+-----------------+
	// |  1 | INSERT      | td    | NULL       | ALL  | NULL          | NULL | NULL    | NULL | NULL |     NULL | NULL            |
	// |  1 | SIMPLE      | td    | NULL       | ALL  | NULL          | NULL | NULL    | NULL |    1 |   100.00 | Using temporary |
	// +----+-------------+-------+------------+------+---------------+------+---------+------+------+----------+-----------------+

	for _, rowAny := range rowList {
		row, ok := rowAny.([]any)
		if !ok {
			return 0, errors.Errorf("expected []any but got %t", row)
		}
		if len(row) != 12 {
			return 0, errors.Errorf("expected 12 but got %d", len(row))
		}
		switch col := row[9].(type) {
		case int:
			return int64(col), nil
		case int32:
			return int64(col), nil
		case int64:
			return col, nil
		case string:
			v, err := strconv.ParseInt(col, 10, 64)
			if err != nil {
				return 0, errors.Errorf("expected int or int64 but got string(%s)", col)
			}
			return v, nil
		default:
			continue
		}
	}

	return 0, errors.Errorf("failed to extract rows from query plan")
}
