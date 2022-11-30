package tables

import (
	"context"
	"database/sql"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
	"github.com/selefra/selefra-provider-snowflake/snowflake_client"

	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-snowflake/table_schema_generator"
)

type TableSnowflakeViewGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableSnowflakeViewGenerator{}

func (x *TableSnowflakeViewGenerator) GetTableName() string {
	return "snowflake_view"
}

func (x *TableSnowflakeViewGenerator) GetTableDescription() string {
	return ""
}

func (x *TableSnowflakeViewGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableSnowflakeViewGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableSnowflakeViewGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {

			db, err := snowflake_client.Connect(ctx, taskClient.(*snowflake_client.Client).Config)
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}
			rows, err := db.QueryContext(ctx, "SHOW VIEWS")
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}
			defer rows.Close()

			for rows.Next() {
				var createdOn sql.NullTime
				var name sql.NullString
				var reserved sql.NullString
				var databaseName sql.NullString
				var schemaName sql.NullString
				var owner sql.NullString
				var comment sql.NullString
				var text sql.NullString
				var isSecure sql.NullString
				var isMaterialized sql.NullString

				err = rows.Scan(&createdOn, &name, &reserved, &databaseName, &schemaName, &owner, &comment, &text, &isSecure, &isMaterialized)
				if err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)
				}
				resultChannel <- View{createdOn, name, reserved, databaseName, schemaName, owner, comment, text, isSecure, isMaterialized}
			}

			for rows.NextResultSet() {
				for rows.Next() {
					var createdOn sql.NullTime
					var name sql.NullString
					var reserved sql.NullString
					var databaseName sql.NullString
					var schemaName sql.NullString
					var owner sql.NullString
					var comment sql.NullString
					var text sql.NullString
					var isSecure sql.NullString
					var isMaterialized sql.NullString
					err = rows.Scan(&createdOn, &name, &reserved, &databaseName, &schemaName, &owner, &comment, &text, &isSecure, &isMaterialized)
					if err != nil {
						return schema.NewDiagnosticsErrorPullTable(task.Table, err)
					}
					resultChannel <- View{createdOn, name, reserved, databaseName, schemaName, owner, comment, text, isSecure, isMaterialized}
				}
			}
			return schema.NewDiagnosticsErrorPullTable(task.Table, nil)

		},
	}
}

type View struct {
	CreatedOn      sql.NullTime   `json:"created_on"`
	Name           sql.NullString `json:"name"`
	Reserved       sql.NullString `json:"reserved"`
	DatabaseName   sql.NullString `json:"database_name"`
	SchemaName     sql.NullString `json:"schema_name"`
	Owner          sql.NullString `json:"owner"`
	Comment        sql.NullString `json:"comment"`
	Text           sql.NullString `json:"text"`
	IsSecure       sql.NullString `json:"is_secure"`
	IsMaterialized sql.NullString `json:"is_materialized"`
}

func (x *TableSnowflakeViewGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return nil
}

func (x *TableSnowflakeViewGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("is_materialized").ColumnType(schema.ColumnTypeBool).Description("True if the view is a materialized view; false otherwise.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {
				return result.(View).IsMaterialized.String, nil
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).Description("The name of the view.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {
				return result.(View).Name.String, nil
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("database_name").ColumnType(schema.ColumnTypeString).Description("The name of the database in which the view exists.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {
				return result.(View).DatabaseName.String, nil
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("created_on").ColumnType(schema.ColumnTypeTimestamp).Description("The timestamp at which the view was created.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {
				return result.(View).CreatedOn.Time, nil
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("owner").ColumnType(schema.ColumnTypeString).Description("The owner of the view.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {
				return result.(View).Owner.String, nil
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("text").ColumnType(schema.ColumnTypeString).Description("The text of the command that created the view, e.g., CREATE VIEW.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {
				return result.(View).Text.String, nil
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("schema_name").ColumnType(schema.ColumnTypeString).Description("The name of the schema in which the view exists.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {
				return result.(View).SchemaName.String, nil
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("comment").ColumnType(schema.ColumnTypeString).Description("Optional comment.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {
				return result.(View).Comment.String, nil
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("is_secure").ColumnType(schema.ColumnTypeBool).Description("True if the view is a secure view; false otherwise.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {
				return result.(View).IsSecure.String, nil
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).Description("The Snowflake region in which the account is located.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {
				return taskClient.(*snowflake_client.Client).Config.Region, nil
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("account").ColumnType(schema.ColumnTypeString).Description("The Snowflake account ID.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {
				return taskClient.(*snowflake_client.Client).Config.Account, nil
			})).Build(),
	}
}

func (x *TableSnowflakeViewGenerator) GetSubTables() []*schema.Table {
	return []*schema.Table{
		table_schema_generator.GenTableSchema(&TableSnowflakeViewGrantGenerator{}),
	}
}
