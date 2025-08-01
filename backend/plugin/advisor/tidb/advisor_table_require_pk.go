package tidb

import (
	"context"
	"fmt"
	"strings"

	"github.com/pingcap/tidb/pkg/parser/ast"
	"github.com/pkg/errors"

	"github.com/bytebase/bytebase/backend/common"
	storepb "github.com/bytebase/bytebase/backend/generated-go/store"
	"github.com/bytebase/bytebase/backend/plugin/advisor"
	"github.com/bytebase/bytebase/backend/plugin/advisor/catalog"
)

const (
	primaryKeyName = "PRIMARY"
)

var (
	_ advisor.Advisor = (*TableRequirePKAdvisor)(nil)
	_ ast.Visitor     = (*tableRequirePKChecker)(nil)
)

func init() {
	advisor.Register(storepb.Engine_TIDB, advisor.MySQLTableRequirePK, &TableRequirePKAdvisor{})
}

// TableRequirePKAdvisor is the advisor checking table requires PK.
type TableRequirePKAdvisor struct {
}

// Check checks table requires PK.
func (*TableRequirePKAdvisor) Check(_ context.Context, checkCtx advisor.Context) ([]*storepb.Advice, error) {
	root, ok := checkCtx.AST.([]ast.StmtNode)
	if !ok {
		return nil, errors.Errorf("failed to convert to StmtNode")
	}

	level, err := advisor.NewStatusBySQLReviewRuleLevel(checkCtx.Rule.Level)
	if err != nil {
		return nil, err
	}
	checker := &tableRequirePKChecker{
		level:   level,
		title:   string(checkCtx.Rule.Type),
		tables:  make(tablePK),
		line:    make(map[string]int),
		catalog: checkCtx.Catalog,
	}

	for _, stmtNode := range root {
		(stmtNode).Accept(checker)
	}

	return checker.generateAdviceList(), nil
}

type tableRequirePKChecker struct {
	adviceList []*storepb.Advice
	level      storepb.Advice_Status
	title      string
	tables     tablePK
	line       map[string]int
	catalog    *catalog.Finder
}

// Enter implements the ast.Visitor interface.
func (v *tableRequirePKChecker) Enter(in ast.Node) (ast.Node, bool) {
	switch node := in.(type) {
	// CREATE TABLE
	case *ast.CreateTableStmt:
		if node.ReferTable != nil {
			v.createTableLike(node)
		} else {
			v.createTable(node)
		}
		v.line[node.Table.Name.O] = node.OriginTextPosition()
	// DROP TABLE
	case *ast.DropTableStmt:
		for _, table := range node.Tables {
			delete(v.tables, table.Name.String())
		}
	// ALTER TABLE
	case *ast.AlterTableStmt:
		tableName := node.Table.Name.O
		for _, spec := range node.Specs {
			switch spec.Tp {
			// ADD CONSTRAINT
			case ast.AlterTableAddConstraint:
				if spec.Constraint.Tp == ast.ConstraintPrimaryKey {
					v.tables[tableName] = newColumnSet(convertConstraintToKeySlice(spec.Constraint))
				}
			// DROP PRIMARY KEY
			case ast.AlterTableDropPrimaryKey:
				v.initEmptyTable(tableName)
				v.line[tableName] = node.OriginTextPosition()
			// DROP INDEX
			case ast.AlterTableDropIndex:
				if strings.ToUpper(spec.Name) == primaryKeyName {
					v.initEmptyTable(tableName)
					v.line[tableName] = node.OriginTextPosition()
				}
			// ADD COLUMNS
			case ast.AlterTableAddColumns:
				v.addPKIfExistByCols(tableName, spec.NewColumns)
			// CHANGE COLUMN
			case ast.AlterTableChangeColumn:
				if v.changeColumn(tableName, spec.OldColumnName.Name.String(), spec.NewColumns[0].Name.Name.String()) {
					v.line[tableName] = node.OriginTextPosition()
				}
				v.addPKIfExistByCols(tableName, spec.NewColumns[:1])
			// MODIFY COLUMN
			case ast.AlterTableModifyColumn:
				v.addPKIfExistByCols(tableName, spec.NewColumns[:1])
			// DROP COLUMN
			case ast.AlterTableDropColumn:
				if v.dropColumn(tableName, spec.OldColumnName.Name.String()) {
					v.line[tableName] = node.OriginTextPosition()
				}
			default:
				// Skip other alter table specification types
			}
		}
	}

	return in, false
}

// Leave implements the ast.Visitor interface.
func (*tableRequirePKChecker) Leave(in ast.Node) (ast.Node, bool) {
	return in, true
}

func (v *tableRequirePKChecker) generateAdviceList() []*storepb.Advice {
	tableList := v.tables.tableList()
	for _, tableName := range tableList {
		if len(v.tables[tableName]) == 0 {
			v.adviceList = append(v.adviceList, &storepb.Advice{
				Status:        v.level,
				Code:          advisor.TableNoPK.Int32(),
				Title:         v.title,
				Content:       fmt.Sprintf("Table `%s` requires PRIMARY KEY", tableName),
				StartPosition: common.ConvertANTLRLineToPosition(v.line[tableName]),
			})
		}
	}

	return v.adviceList
}

func (v *tableRequirePKChecker) initEmptyTable(name string) {
	v.tables[name] = make(columnSet)
}

func (v *tableRequirePKChecker) createTable(node *ast.CreateTableStmt) {
	table := node.Table.Name.String()
	v.initEmptyTable(table)
	v.addPKIfExistByCols(table, node.Cols)

	for _, constraint := range node.Constraints {
		if constraint.Tp == ast.ConstraintPrimaryKey {
			v.tables[table] = newColumnSet(convertConstraintToKeySlice(constraint))
		}
	}
}

func (v *tableRequirePKChecker) createTableLike(node *ast.CreateTableStmt) {
	table := node.Table.Name.String()
	v.initEmptyTable(table)
	referTableName := node.ReferTable.Name.String()
	// Switch to check refer table's primary key.
	// Find refer table in context first, and then find it in catalog.
	if referTablePk, ok := v.tables[referTableName]; ok {
		var columns []string
		for column := range referTablePk {
			columns = append(columns, column)
		}
		v.tables[table] = newColumnSet(columns)
	} else {
		referTableState := v.catalog.Origin.FindTable(&catalog.TableFind{
			TableName: referTableName,
		})
		if referTableState != nil {
			indexSet := referTableState.Index(&catalog.TableIndexFind{})
			for _, index := range *indexSet {
				if index.Primary() {
					v.tables[table] = newColumnSet(index.ExpressionList())
				}
			}
		}
	}
}

func (v *tableRequirePKChecker) dropColumn(table string, column string) bool {
	if _, ok := v.tables[table]; !ok {
		_, pk := v.catalog.Origin.FindIndex(&catalog.IndexFind{
			TableName: table,
			IndexName: primaryKeyName,
		})
		if pk == nil {
			return false
		}
		v.tables[table] = newColumnSet(pk.ExpressionList())
	}

	pk := v.tables[table]
	_, columnInPK := pk[column]
	delete(v.tables[table], column)
	return columnInPK
}

func (v *tableRequirePKChecker) changeColumn(table string, oldColumn string, newColumn string) bool {
	if v.dropColumn(table, oldColumn) {
		pk := v.tables[table]
		pk[newColumn] = true
		return true
	}
	return false
}

func (v *tableRequirePKChecker) addPKIfExistByCols(table string, columns []*ast.ColumnDef) {
	for _, column := range columns {
		for _, option := range column.Options {
			if option.Tp == ast.ColumnOptionPrimaryKey {
				v.tables[table] = newColumnSet([]string{column.Name.Name.String()})
				return
			}
		}
	}
}

func convertConstraintToKeySlice(constraint *ast.Constraint) []string {
	var columns []string
	for _, key := range constraint.Keys {
		columns = append(columns, key.Column.Name.String())
	}
	return columns
}
