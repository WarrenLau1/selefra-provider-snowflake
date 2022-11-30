package tables

import (
	"context"
	"database/sql"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
	"github.com/selefra/selefra-provider-snowflake/snowflake_client"

	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-snowflake/table_schema_generator"
)

type TableSnowflakeSchemataGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableSnowflakeSchemataGenerator{}

func (x *TableSnowflakeSchemataGenerator) GetTableName() string {
	return "snowflake_schemata"
}

func (x *TableSnowflakeSchemataGenerator) GetTableDescription() string {
	return ""
}

func (x *TableSnowflakeSchemataGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableSnowflakeSchemataGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableSnowflakeSchemataGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {

			db, err := snowflake_client.Connect(ctx, taskClient.(*snowflake_client.Client).Config)
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}
			rows, err := db.QueryContext(ctx, "select * from SNOWFLAKE.ACCOUNT_USAGE.schemata;")
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}
			defer rows.Close()

			columns, err := rows.Columns()
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}

			for rows.Next() {
				s := Schemata{}

				cols := make([]interface{}, len(columns))
				for i, col := range columns {
					cols[i] = SchemataCol(col, &s)
				}

				err = rows.Scan(cols...)
				if err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)
				}
				resultChannel <- s
			}

			for rows.NextResultSet() {
				for rows.Next() {
					s := Schemata{}

					cols := make([]interface{}, len(columns))
					for i, col := range columns {
						cols[i] = SchemataCol(col, &s)
					}

					err = rows.Scan(cols...)
					if err != nil {
						return schema.NewDiagnosticsErrorPullTable(task.Table, err)
					}
					resultChannel <- s
				}
			}
			return schema.NewDiagnosticsErrorPullTable(task.Table, nil)

		},
	}
}

type Schemata struct {
	SchemaId                   sql.NullInt64  `json:"SCHEMA_ID"`
	SchemaName                 sql.NullString `json:"SCHEMA_NAME"`
	CatalogId                  sql.NullInt64  `json:"CATALOG_ID"`
	CatalogName                sql.NullString `json:"CATALOG_NAME"`
	SchemaOwner                sql.NullString `json:"SCHEMA_OWNER"`
	RetentionTime              sql.NullInt64  `json:"RETENTION_TIME"`
	IsTransient                sql.NullString `json:"IS_TRANSIENT"`
	IsManagedAccess            sql.NullString `json:"IS_MANAGED_ACCESS"`
	DefaultCharacterSetCatalog sql.NullString `json:"DEFAULT_CHARACTER_SET_CATALOG"`
	DefaultCharacterSetSchema  sql.NullString `json:"DEFAULT_CHARACTER_SET_SCHEMA"`
	DefaultCharacterSetName    sql.NullString `json:"DEFAULT_CHARACTER_SET_NAME"`
	SqlPath                    sql.NullString `json:"SQL_PATH"`
	Comment                    sql.NullString `json:"COMMENT"`
	Created                    sql.NullTime   `json:"CREATED"`
	LastAltered                sql.NullTime   `json:"LAST_ALTERED"`
	Deleted                    sql.NullTime   `json:"DELETED"`
}

// SchemataCol returns a reference for a column of a Schemata
func SchemataCol(colname string, item *Schemata) interface{} {
	switch colname {
	case "SCHEMA_ID":
		return &item.SchemaId
	case "SCHEMA_NAME":
		return &item.SchemaName
	case "CATALOG_ID":
		return &item.CatalogId
	case "CATALOG_NAME":
		return &item.CatalogName
	case "SCHEMA_OWNER":
		return &item.SchemaOwner
	case "RETENTION_TIME":
		return &item.RetentionTime
	case "IS_TRANSIENT":
		return &item.IsTransient
	case "IS_MANAGED_ACCESS":
		return &item.IsManagedAccess
	case "DEFAULT_CHARACTER_SET_CATALOG":
		return &item.DefaultCharacterSetCatalog
	case "DEFAULT_CHARACTER_SET_SCHEMA":
		return &item.DefaultCharacterSetSchema
	case "DEFAULT_CHARACTER_SET_NAME":
		return &item.DefaultCharacterSetName
	case "SQL_PATH":
		return &item.SqlPath
	case "COMMENT":
		return &item.Comment
	case "CREATED":
		return &item.Created
	case "LAST_ALTERED":
		return &item.LastAltered
	case "DELETED":
		return &item.Deleted
	default:
		panic("unknown column " + colname)
	}
}

func (x *TableSnowflakeSchemataGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return nil
}

func (x *TableSnowflakeSchemataGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).Description("The Snowflake region in which the account is located.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {
				return taskClient.(*snowflake_client.Client).Config.Region, nil
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("is_transient").ColumnType(schema.ColumnTypeString).Description("Whether this is a transient schema.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {
				return result.(Schemata).IsTransient.String, nil
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("comment").ColumnType(schema.ColumnTypeString).Description("Comment for this schema.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {
				return result.(Schemata).Comment.String, nil
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("created").ColumnType(schema.ColumnTypeTimestamp).Description("Creation time of the schema.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {
				return result.(Schemata).Created.Time, nil
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("schema_id").ColumnType(schema.ColumnTypeString).Description("ID of the schema.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {
				return result.(Schemata).SchemaId.Int64, nil
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("schema_owner").ColumnType(schema.ColumnTypeString).Description("Name of the role that owns the schema.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {
				return result.(Schemata).SchemaOwner.String, nil
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("deleted").ColumnType(schema.ColumnTypeTimestamp).Description("Deletion time of the schema.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {
				return result.(Schemata).Deleted.Time, nil
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("schema_name").ColumnType(schema.ColumnTypeString).Description("Name of the schema.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {
				return result.(Schemata).SchemaName.String, nil
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("catalog_name").ColumnType(schema.ColumnTypeString).Description("Database that the schema belongs to.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {
				return result.(Schemata).CatalogName.String, nil
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("retention_time").ColumnType(schema.ColumnTypeInt).Description("Number of days that historical data is retained for Time Travel.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {
				return result.(Schemata).RetentionTime.Int64, nil
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("is_managed_access").ColumnType(schema.ColumnTypeString).Description("Whether the schema is a managed access schema.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {
				return result.(Schemata).IsManagedAccess.String, nil
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("last_altered").ColumnType(schema.ColumnTypeTimestamp).Description("Last altered time of the schema.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {
				return result.(Schemata).LastAltered.Time, nil
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("account").ColumnType(schema.ColumnTypeString).Description("The Snowflake account ID.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {
				return taskClient.(*snowflake_client.Client).Config.Account, nil
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("catalog_id").ColumnType(schema.ColumnTypeString).Description("ID of the database that the schema belongs to.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {
				return result.(Schemata).CatalogId.Int64, nil
			})).Build(),
	}
}

func (x *TableSnowflakeSchemataGenerator) GetSubTables() []*schema.Table {
	return nil
}
