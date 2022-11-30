package tables

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
	"github.com/selefra/selefra-provider-snowflake/snowflake_client"

	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-snowflake/table_schema_generator"
)

type TableSnowflakeUserGrantGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableSnowflakeUserGrantGenerator{}

func (x *TableSnowflakeUserGrantGenerator) GetTableName() string {
	return "snowflake_user_grant"
}

func (x *TableSnowflakeUserGrantGenerator) GetTableDescription() string {
	return ""
}

func (x *TableSnowflakeUserGrantGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableSnowflakeUserGrantGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableSnowflakeUserGrantGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {

			user := task.ParentRawResult.(User).Name.String

			if user == "" {
				return schema.NewDiagnosticsErrorPullTable(task.Table, nil)
			}
			db, err := snowflake_client.Connect(ctx, taskClient.(*snowflake_client.Client).Config)
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}

			rows, err := db.QueryContext(ctx, fmt.Sprintf("SHOW GRANTS TO USER \"%s\"", user))
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}
			defer rows.Close()

			columns, err := rows.Columns()
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}

			for rows.Next() {
				userGrant := UserGrant{}

				cols := make([]interface{}, len(columns))

				for i, col := range columns {
					cols[i] = UserGrantCol(col, &userGrant)
				}

				err = rows.Scan(cols...)
				if err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)
				}
				resultChannel <- userGrant
			}

			for rows.NextResultSet() {
				for rows.Next() {
					userGrant := UserGrant{}

					cols := make([]interface{}, len(columns))

					for i, col := range columns {
						cols[i] = UserGrantCol(col, &userGrant)
					}

					err = rows.Scan(cols...)
					if err != nil {
						return schema.NewDiagnosticsErrorPullTable(task.Table, err)
					}
					resultChannel <- userGrant
				}
			}
			return schema.NewDiagnosticsErrorPullTable(task.Table, nil)

		},
	}
}

type RoleGrant struct {
	CreatedOn   sql.NullTime   `json:"created_on"`
	Role        sql.NullString `json:"role"`
	GrantedTo   sql.NullString `json:"granted_to"`
	GranteeName sql.NullString `json:"grantee_name"`
	GrantedBy   sql.NullString `json:"granted_by"`
}
type UserGrant RoleGrant

// UserGrantCol returns a reference for a column of a UserGrant
func UserGrantCol(colname string, sp *UserGrant) interface{} {
	switch colname {
	case "created_on":
		return &sp.CreatedOn
	case "role":
		return &sp.Role
	case "granted_to":
		return &sp.GrantedTo
	case "grantee_name":
		return &sp.GranteeName
	case "granted_by":
		return &sp.GrantedBy
	default:
		panic("unknown column " + colname)
	}
}

func (x *TableSnowflakeUserGrantGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return nil
}

func (x *TableSnowflakeUserGrantGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("account").ColumnType(schema.ColumnTypeString).Description("The Snowflake account ID.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {
				return taskClient.(*snowflake_client.Client).Config.Account, nil
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("role").ColumnType(schema.ColumnTypeString).Description("Name of the role that has been granted to user.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {
				return result.(UserGrant).Role.String, nil
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("created_on").ColumnType(schema.ColumnTypeTimestamp).Description("Date and time when the role was granted to the user/role.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {
				return result.(UserGrant).CreatedOn.Time, nil
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("granted_to").ColumnType(schema.ColumnTypeString).Description("Type of the object. Only USER for this table.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {
				return result.(UserGrant).GrantedTo.String, nil
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("granted_by").ColumnType(schema.ColumnTypeString).Description("Name of the object that granted access on the user.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {
				return result.(UserGrant).GrantedBy.String, nil
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).Description("The Snowflake region in which the account is located.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {
				return taskClient.(*snowflake_client.Client).Config.Region, nil
			})).Build(),
	}
}

func (x *TableSnowflakeUserGrantGenerator) GetSubTables() []*schema.Table {
	return nil
}
