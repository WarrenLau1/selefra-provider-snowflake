package tables

import (
	"context"
	"database/sql"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
	"github.com/selefra/selefra-provider-snowflake/snowflake_client"

	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-snowflake/table_schema_generator"
)

type TableSnowflakeDatabaseGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableSnowflakeDatabaseGenerator{}

func (x *TableSnowflakeDatabaseGenerator) GetTableName() string {
	return "snowflake_database"
}

func (x *TableSnowflakeDatabaseGenerator) GetTableDescription() string {
	return ""
}

func (x *TableSnowflakeDatabaseGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableSnowflakeDatabaseGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableSnowflakeDatabaseGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {

			db, err := snowflake_client.Connect(ctx, taskClient.(*snowflake_client.Client).Config)
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}
			rows, err := db.QueryContext(ctx, "SHOW DATABASES")
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}
			defer rows.Close()

			for rows.Next() {
				var CreatedOn sql.NullString
				var Name sql.NullString
				var IsDefault sql.NullString
				var IsCurrent sql.NullString
				var Origin sql.NullString
				var Owner sql.NullString
				var Comment sql.NullString
				var Options sql.NullString
				var RetentionTime sql.NullString

				err = rows.Scan(&CreatedOn, &Name, &IsDefault, &IsCurrent, &Origin, &Owner, &Comment, &Options, &RetentionTime)
				if err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)
				}
				resultChannel <- Database{CreatedOn, Name, IsDefault, IsCurrent, Origin, Owner, Comment, Options, RetentionTime}
			}

			for rows.NextResultSet() {
				for rows.Next() {
					var CreatedOn sql.NullString
					var Name sql.NullString
					var IsDefault sql.NullString
					var IsCurrent sql.NullString
					var Origin sql.NullString
					var Owner sql.NullString
					var Comment sql.NullString
					var Options sql.NullString
					var RetentionTime sql.NullString

					err = rows.Scan(&CreatedOn, &Name, &IsDefault, &IsCurrent, &Origin, &Owner, &Comment, &Options, &RetentionTime)
					if err != nil {
						return schema.NewDiagnosticsErrorPullTable(task.Table, err)
					}
					resultChannel <- Database{CreatedOn, Name, IsDefault, IsCurrent, Origin, Owner, Comment, Options, RetentionTime}
				}
			}
			return schema.NewDiagnosticsErrorPullTable(task.Table, nil)

		},
	}
}

type Database struct {
	CreatedOn     sql.NullString `json:"created_on"`
	Name          sql.NullString `json:"name"`
	IsDefault     sql.NullString `json:"is_default"`
	IsCurrent     sql.NullString `json:"is_current"`
	Origin        sql.NullString `json:"origin"`
	Owner         sql.NullString `json:"owner"`
	Comment       sql.NullString `json:"comment"`
	Options       sql.NullString `json:"options"`
	RetentionTime sql.NullString `json:"retention_time"`
}

func (x *TableSnowflakeDatabaseGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return nil
}

func (x *TableSnowflakeDatabaseGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("account").ColumnType(schema.ColumnTypeString).Description("The Snowflake account ID.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {
				return taskClient.(*snowflake_client.Client).Config.Account, nil
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).Description("Name of the database.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {
				return result.(Database).Name.String, nil
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("created_on").ColumnType(schema.ColumnTypeTimestamp).Description("Creation time of the database.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {
				return result.(Database).CreatedOn.String, nil
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("is_default").ColumnType(schema.ColumnTypeString).Description("Name of the default database for authenticating user.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {
				return result.(Database).IsDefault.String, nil
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("origin").ColumnType(schema.ColumnTypeString).Description("Name of the origin database.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {
				return result.(Database).Origin.String, nil
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).Description("The Snowflake region in which the account is located.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {
				return taskClient.(*snowflake_client.Client).Config.Region, nil
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("comment").ColumnType(schema.ColumnTypeString).Description("Comment for this database.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {
				return result.(Database).Comment.String, nil
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("is_current").ColumnType(schema.ColumnTypeString).Description("Name of the current database for authenticating user.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {
				return result.(Database).IsCurrent.String, nil
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("options").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {
				return result.(Database).Options.String, nil
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("owner").ColumnType(schema.ColumnTypeString).Description("Name of the role that owns the schema.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {
				return result.(Database).Owner.String, nil
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("retention_time").ColumnType(schema.ColumnTypeInt).Description("Number of days that historical data is retained for Time Travel.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {
				return result.(Database).RetentionTime.String, nil
			})).Build(),
	}
}

func (x *TableSnowflakeDatabaseGenerator) GetSubTables() []*schema.Table {
	return []*schema.Table{
		table_schema_generator.GenTableSchema(&TableSnowflakeDatabaseGrantGenerator{}),
	}
}
