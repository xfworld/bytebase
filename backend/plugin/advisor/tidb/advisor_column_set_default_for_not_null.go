package tidb

// Framework code is generated by the generator.

import (
	"context"
	"fmt"

	"github.com/pingcap/tidb/pkg/parser/ast"
	"github.com/pkg/errors"

	"github.com/bytebase/bytebase/backend/common"
	storepb "github.com/bytebase/bytebase/backend/generated-go/store"
	"github.com/bytebase/bytebase/backend/plugin/advisor"
)

var (
	_ advisor.Advisor = (*ColumnSetDefaultForNotNullAdvisor)(nil)
	_ ast.Visitor     = (*columnSetDefaultForNotNullChecker)(nil)
)

func init() {
	advisor.Register(storepb.Engine_TIDB, advisor.MySQLColumnSetDefaultForNotNull, &ColumnSetDefaultForNotNullAdvisor{})
}

// ColumnSetDefaultForNotNullAdvisor is the advisor checking for set default value for not null column.
type ColumnSetDefaultForNotNullAdvisor struct {
}

// Check checks for set default value for not null column.
func (*ColumnSetDefaultForNotNullAdvisor) Check(_ context.Context, checkCtx advisor.Context) ([]*storepb.Advice, error) {
	stmtList, ok := checkCtx.AST.([]ast.StmtNode)
	if !ok {
		return nil, errors.Errorf("failed to convert to StmtNode")
	}

	level, err := advisor.NewStatusBySQLReviewRuleLevel(checkCtx.Rule.Level)
	if err != nil {
		return nil, err
	}
	checker := &columnSetDefaultForNotNullChecker{
		level: level,
		title: string(checkCtx.Rule.Type),
	}

	for _, stmt := range stmtList {
		checker.text = stmt.Text()
		(stmt).Accept(checker)
	}

	return checker.adviceList, nil
}

type columnSetDefaultForNotNullChecker struct {
	adviceList []*storepb.Advice
	level      storepb.Advice_Status
	title      string
	text       string
}

// Enter implements the ast.Visitor interface.
func (checker *columnSetDefaultForNotNullChecker) Enter(in ast.Node) (ast.Node, bool) {
	var notNullColumnWithNoDefault []columnName
	switch node := in.(type) {
	// CREATE TABLE
	case *ast.CreateTableStmt:
		pkColumn := make(map[string]bool)
		for _, cons := range node.Constraints {
			if cons.Tp == ast.ConstraintPrimaryKey {
				for _, key := range cons.Keys {
					pkColumn[key.Column.Name.O] = true
				}
			}
		}

		for _, column := range node.Cols {
			isPk := pkColumn[column.Name.Name.O]
			if isPk {
				continue
			}
			if !canNull(column) && !setDefault(column) && needDefault(column) {
				notNullColumnWithNoDefault = append(notNullColumnWithNoDefault, columnName{
					tableName:  node.Table.Name.O,
					columnName: column.Name.Name.O,
					line:       column.OriginTextPosition(),
				})
			}
		}
	// ALTER TABLE
	case *ast.AlterTableStmt:
		for _, spec := range node.Specs {
			switch spec.Tp {
			// ADD COLUMNS
			case ast.AlterTableAddColumns:
				for _, column := range spec.NewColumns {
					if !canNull(column) && !setDefault(column) && needDefault(column) {
						notNullColumnWithNoDefault = append(notNullColumnWithNoDefault, columnName{
							tableName:  node.Table.Name.O,
							columnName: column.Name.Name.O,
							line:       node.OriginTextPosition(),
						})
					}
				}
			// CHANGE COLUMN and MODIFY COLUMN
			case ast.AlterTableChangeColumn, ast.AlterTableModifyColumn:
				if !canNull(spec.NewColumns[0]) && !setDefault(spec.NewColumns[0]) && needDefault(spec.NewColumns[0]) {
					notNullColumnWithNoDefault = append(notNullColumnWithNoDefault, columnName{
						tableName:  node.Table.Name.O,
						columnName: spec.NewColumns[0].Name.Name.O,
						line:       node.OriginTextPosition(),
					})
				}
			default:
				// Skip other alter table specification types
			}
		}
	}

	for _, column := range notNullColumnWithNoDefault {
		checker.adviceList = append(checker.adviceList, &storepb.Advice{
			Status:        checker.level,
			Code:          advisor.NotNullColumnWithNoDefault.Int32(),
			Title:         checker.title,
			Content:       fmt.Sprintf("Column `%s`.`%s` is NOT NULL but doesn't have DEFAULT", column.tableName, column.columnName),
			StartPosition: common.ConvertANTLRLineToPosition(column.line),
		})
	}

	return in, false
}

// Leave implements the ast.Visitor interface.
func (*columnSetDefaultForNotNullChecker) Leave(in ast.Node) (ast.Node, bool) {
	return in, true
}

func setDefault(column *ast.ColumnDef) bool {
	for _, option := range column.Options {
		if option.Tp == ast.ColumnOptionDefaultValue {
			return true
		}
	}
	return false
}
