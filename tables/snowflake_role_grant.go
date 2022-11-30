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

type TableSnowflakeRoleGrantGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableSnowflakeRoleGrantGenerator{}

func (x *TableSnowflakeRoleGrantGenerator) GetTableName() string {
	return "snowflake_role_grant"
}

func (x *TableSnowflakeRoleGrantGenerator) GetTableDescription() string {
	return ""
}

func (x *TableSnowflakeRoleGrantGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableSnowflakeRoleGrantGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableSnowflakeRoleGrantGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {

			role := task.ParentRawResult.(Role).Name.String

			if role == "" {
				return schema.NewDiagnosticsErrorPullTable(task.Table, nil)
			}
			db, err := snowflake_client.Connect(ctx, taskClient.(*snowflake_client.Client).Config)
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}
			rows, err := db.QueryContext(ctx, fmt.Sprintf("SHOW GRANTS OF ROLE %s", role))
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}
			defer rows.Close()

			for rows.Next() {
				var CreatedOn sql.NullTime
				var Role sql.NullString
				var GrantedTo sql.NullString
				var GranteeName sql.NullString
				var GrantedBy sql.NullString

				err = rows.Scan(&CreatedOn, &Role, &GrantedTo, &GranteeName, &GrantedBy)
				if err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)
				}
				resultChannel <- RoleGrant{CreatedOn, Role, GrantedTo, GranteeName, GrantedBy}
			}

			for rows.NextResultSet() {
				var CreatedOn sql.NullTime
				var Role sql.NullString
				var GrantedTo sql.NullString
				var GranteeName sql.NullString
				var GrantedBy sql.NullString

				err = rows.Scan(&CreatedOn, &Role, &GrantedTo, &GranteeName, &GrantedBy)
				if err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)
				}
				resultChannel <- RoleGrant{CreatedOn, Role, GrantedTo, GranteeName, GrantedBy}
			}

			return schema.NewDiagnosticsErrorPullTable(task.Table, nil)

		},
	}
}

func (x *TableSnowflakeRoleGrantGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return nil
}

func (x *TableSnowflakeRoleGrantGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("role").ColumnType(schema.ColumnTypeString).Description("Name of the role on that access has been granted.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {
				return result.(RoleGrant).Role.String, nil
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("created_on").ColumnType(schema.ColumnTypeTimestamp).Description("Date and time when the role was granted to the user/role.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {
				return result.(RoleGrant).CreatedOn.Time, nil
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("granted_to").ColumnType(schema.ColumnTypeString).Description("Type of the object. Valid values USER and ROLE.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {
				return result.(RoleGrant).GrantedTo.String, nil
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("grantee_name").ColumnType(schema.ColumnTypeString).Description("Name of the object role has been granted.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {
				return result.(RoleGrant).GranteeName.String, nil
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("granted_by").ColumnType(schema.ColumnTypeString).Description("Name of the object that granted access on the role.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {
				return result.(RoleGrant).GrantedBy.String, nil
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

func (x *TableSnowflakeRoleGrantGenerator) GetSubTables() []*schema.Table {
	return nil
}
