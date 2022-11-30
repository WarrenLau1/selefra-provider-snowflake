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

type TableSnowflakeViewGrantGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableSnowflakeViewGrantGenerator{}

func (x *TableSnowflakeViewGrantGenerator) GetTableName() string {
	return "snowflake_view_grant"
}

func (x *TableSnowflakeViewGrantGenerator) GetTableDescription() string {
	return ""
}

func (x *TableSnowflakeViewGrantGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableSnowflakeViewGrantGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableSnowflakeViewGrantGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {

			view := task.ParentRawResult.(View).Name.String

			if view == "" {
				return schema.NewDiagnosticsErrorPullTable(task.Table, nil)
			}
			database := task.ParentRawResult.(View).DatabaseName.String
			if database == "" {
				return schema.NewDiagnosticsErrorPullTable(task.Table, nil)
			}
			s := task.ParentRawResult.(View).SchemaName.String
			if s == "" {
				return schema.NewDiagnosticsErrorPullTable(task.Table, nil)
			}
			db, err := snowflake_client.Connect(ctx, taskClient.(*snowflake_client.Client).Config)
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}
			rows, err := db.QueryContext(ctx, fmt.Sprintf("SHOW GRANTS ON VIEW %s.%s.%s", database, s, view))
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}
			defer rows.Close()

			for rows.Next() {
				var createdOn sql.NullTime
				var privilege sql.NullString
				var grantedOn sql.NullString
				var name sql.NullString
				var grantedTo sql.NullString
				var granteeName sql.NullString
				var grantOption sql.NullString
				var grantedBy sql.NullString

				err = rows.Scan(&createdOn, &privilege, &grantedOn, &name, &grantedTo, &granteeName, &grantOption, &grantedBy)
				if err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)
				}
				resultChannel <- ViewGrant{createdOn, privilege, grantedOn, name, grantedTo, granteeName, grantOption, grantedBy}
			}

			for rows.NextResultSet() {
				for rows.Next() {
					var createdOn sql.NullTime
					var privilege sql.NullString
					var grantedOn sql.NullString
					var name sql.NullString
					var grantedTo sql.NullString
					var granteeName sql.NullString
					var grantOption sql.NullString
					var grantedBy sql.NullString

					err = rows.Scan(&createdOn, &privilege, &grantedOn, &name, &grantedTo, &granteeName, &grantOption, &grantedBy)
					if err != nil {
						return schema.NewDiagnosticsErrorPullTable(task.Table, err)
					}
					resultChannel <- ViewGrant{createdOn, privilege, grantedOn, name, grantedTo, granteeName, grantOption, grantedBy}
				}
			}
			return schema.NewDiagnosticsErrorPullTable(task.Table, nil)

		},
	}
}

type ViewGrant AccountGrant

func (x *TableSnowflakeViewGrantGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return nil
}

func (x *TableSnowflakeViewGrantGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("grantee_name").ColumnType(schema.ColumnTypeString).Description("Name of the object role has been granted.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {
				return result.(ViewGrant).GranteeName.String, nil
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).Description("The Snowflake region in which the account is located.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {
				return taskClient.(*snowflake_client.Client).Config.Region, nil
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("privilege").ColumnType(schema.ColumnTypeString).Description("A defined level of access to an object.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {
				return result.(ViewGrant).Privilege.String, nil
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("created_on").ColumnType(schema.ColumnTypeTimestamp).Description("Date and time privilege was granted.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {
				return result.(ViewGrant).CreatedOn.Time, nil
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("grant_option").ColumnType(schema.ColumnTypeBool).Description("If set to TRUE, the recipient role can grant the privilege to other roles.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {
				return result.(ViewGrant).GrantOption.String, nil
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("granted_by").ColumnType(schema.ColumnTypeString).Description("Name of the object that granted access on the role.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {
				return result.(ViewGrant).GrantedBy.String, nil
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("granted_on").ColumnType(schema.ColumnTypeString).Description("Date and time when the access was granted.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {
				return result.(ViewGrant).GrantedOn.String, nil
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("granted_to").ColumnType(schema.ColumnTypeString).Description("Type of the object.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {
				return result.(ViewGrant).GrantedTo.String, nil
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("account").ColumnType(schema.ColumnTypeString).Description("The Snowflake account ID.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {
				return taskClient.(*snowflake_client.Client).Config.Account, nil
			})).Build(),
	}
}

func (x *TableSnowflakeViewGrantGenerator) GetSubTables() []*schema.Table {
	return nil
}
