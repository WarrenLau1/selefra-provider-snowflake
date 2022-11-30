package tables

import (
	"context"
	"database/sql"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
	"github.com/selefra/selefra-provider-snowflake/snowflake_client"

	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-snowflake/table_schema_generator"
)

type TableSnowflakeAccountParameterGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableSnowflakeAccountParameterGenerator{}

func (x *TableSnowflakeAccountParameterGenerator) GetTableName() string {
	return "snowflake_account_parameter"
}

func (x *TableSnowflakeAccountParameterGenerator) GetTableDescription() string {
	return ""
}

func (x *TableSnowflakeAccountParameterGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableSnowflakeAccountParameterGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableSnowflakeAccountParameterGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {

			db, err := snowflake_client.Connect(ctx, taskClient.(*snowflake_client.Client).Config)
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}
			rows, err := db.QueryContext(ctx, "SHOW PARAMETERS IN ACCOUNT;")
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}
			defer rows.Close()

			columns, err := rows.Columns()
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}

			for rows.Next() {
				parameter := Parameter{}

				cols := make([]interface{}, len(columns))
				for i, col := range columns {
					cols[i] = ParameterCol(col, &parameter)
				}

				err = rows.Scan(cols...)
				if err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)
				}
				resultChannel <- parameter
			}

			for rows.NextResultSet() {
				for rows.Next() {
					parameter := Parameter{}

					cols := make([]interface{}, len(columns))
					for i, col := range columns {
						cols[i] = ParameterCol(col, &parameter)
					}

					err = rows.Scan(cols...)
					if err != nil {
						return schema.NewDiagnosticsErrorPullTable(task.Table, err)
					}
					resultChannel <- parameter
				}
			}
			return schema.NewDiagnosticsErrorPullTable(task.Table, nil)

		},
	}
}

type Parameter struct {
	Key         sql.NullString `json:"key"`
	Value       sql.NullString `json:"value"`
	Default     sql.NullString `json:"default"`
	Level       sql.NullString `json:"level"`
	Description sql.NullString `json:"description"`
	Type        sql.NullString `json:"type"`
}

// ParameterCol returns a reference for a column of a Parameter
func ParameterCol(colname string, item *Parameter) interface{} {
	switch colname {
	case "key":
		return &item.Key
	case "value":
		return &item.Value
	case "default":
		return &item.Default
	case "level":
		return &item.Level
	case "description":
		return &item.Description
	case "type":
		return &item.Type
	default:
		panic("unknown column " + colname)
	}
}

func (x *TableSnowflakeAccountParameterGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return nil
}

func (x *TableSnowflakeAccountParameterGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("default").ColumnType(schema.ColumnTypeString).Description("Default value of the parameter.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {
				return result.(Parameter).Default.String, nil
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("description").ColumnType(schema.ColumnTypeString).Description("Description for the parameter.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {
				return result.(Parameter).Description.String, nil
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("level").ColumnType(schema.ColumnTypeString).Description("Level of the parameter. Can be SYSTEM or ACCOUNT.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {
				return result.(Parameter).Level.String, nil
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("type").ColumnType(schema.ColumnTypeString).Description("Data type of the parameter value.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {
				return result.(Parameter).Type.String, nil
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).Description("The Snowflake region in which the account is located.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {
				return taskClient.(*snowflake_client.Client).Config.Region, nil
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("account").ColumnType(schema.ColumnTypeString).Description("The Snowflake account ID.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {
				return taskClient.(*snowflake_client.Client).Config.Account, nil
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("key").ColumnType(schema.ColumnTypeString).Description("Name of the account parameter.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {
				return result.(Parameter).Key.String, nil
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("value").ColumnType(schema.ColumnTypeString).Description("Current value of the parameter.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {
				return result.(Parameter).Value.String, nil
			})).Build(),
	}
}

func (x *TableSnowflakeAccountParameterGenerator) GetSubTables() []*schema.Table {
	return nil
}
