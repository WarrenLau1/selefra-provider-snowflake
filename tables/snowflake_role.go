package tables

import (
	"context"
	"database/sql"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
	"github.com/selefra/selefra-provider-snowflake/snowflake_client"

	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-snowflake/table_schema_generator"
)

type TableSnowflakeRoleGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableSnowflakeRoleGenerator{}

func (x *TableSnowflakeRoleGenerator) GetTableName() string {
	return "snowflake_role"
}

func (x *TableSnowflakeRoleGenerator) GetTableDescription() string {
	return ""
}

func (x *TableSnowflakeRoleGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableSnowflakeRoleGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableSnowflakeRoleGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {

			db, err := snowflake_client.Connect(ctx, taskClient.(*snowflake_client.Client).Config)
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}
			rows, err := db.QueryContext(ctx, "SHOW ROLES")
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}
			defer rows.Close()

			for rows.Next() {
				var CreatedOn sql.NullTime
				var Name sql.NullString
				var IsDefault sql.NullString
				var IsCurrent sql.NullString
				var IsInherited sql.NullString
				var AssignedToUsers sql.NullInt64
				var GrantedToRoles sql.NullInt64
				var GrantedRoles sql.NullInt64
				var Owner sql.NullString
				var Comment sql.NullString

				err = rows.Scan(&CreatedOn, &Name, &IsDefault, &IsCurrent, &IsInherited, &AssignedToUsers, &GrantedToRoles, &GrantedRoles, &Owner, &Comment)
				if err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)
				}
				resultChannel <- Role{CreatedOn, Name, IsDefault, IsCurrent, IsInherited, AssignedToUsers, GrantedToRoles, GrantedRoles, Owner, Comment}
			}

			for rows.NextResultSet() {
				for rows.Next() {
					var CreatedOn sql.NullTime
					var Name sql.NullString
					var IsDefault sql.NullString
					var IsCurrent sql.NullString
					var IsInherited sql.NullString
					var AssignedToUsers sql.NullInt64
					var GrantedToRoles sql.NullInt64
					var GrantedRoles sql.NullInt64
					var Owner sql.NullString
					var Comment sql.NullString

					err = rows.Scan(&CreatedOn, &Name, &IsDefault, &IsCurrent, &IsInherited, &AssignedToUsers, &GrantedToRoles, &GrantedRoles, &Owner, &Comment)
					if err != nil {
						return schema.NewDiagnosticsErrorPullTable(task.Table, err)
					}
					resultChannel <- Role{CreatedOn, Name, IsDefault, IsCurrent, IsInherited, AssignedToUsers, GrantedToRoles, GrantedRoles, Owner, Comment}
				}
			}
			return schema.NewDiagnosticsErrorPullTable(task.Table, nil)

		},
	}
}

func (x *TableSnowflakeRoleGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return nil
}

func (x *TableSnowflakeRoleGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("comment").ColumnType(schema.ColumnTypeString).Description("Comment for the role.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {
				return result.(Role).Comment.String, nil
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("granted_to_roles").ColumnType(schema.ColumnTypeInt).Description("Number of roles that inherit the privileges of this role.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {
				return result.(Role).GrantedToRoles.Int64, nil
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("is_default").ColumnType(schema.ColumnTypeString).Description("\"Y\" if is the default role of authenticated user, otherwise \"F\".").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {
				return result.(Role).IsDefault.String, nil
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("is_inherited").ColumnType(schema.ColumnTypeString).Description("\"Y\" if current role is inherited by authenticated user, otherwise \"F\".").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {
				return result.(Role).IsInherited.String, nil
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("owner").ColumnType(schema.ColumnTypeString).Description("Owner of the role.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {
				return result.(Role).Owner.String, nil
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("account").ColumnType(schema.ColumnTypeString).Description("The Snowflake account ID.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {
				return taskClient.(*snowflake_client.Client).Config.Account, nil
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).Description("Name of the role.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {
				return result.(Role).Name.String, nil
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("assigned_to_users").ColumnType(schema.ColumnTypeInt).Description("Number of users the role is assigned.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {
				return result.(Role).AssignedToUsers.Int64, nil
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("granted_roles").ColumnType(schema.ColumnTypeInt).Description("Number of roles inherited by this role.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {
				return result.(Role).GrantedRoles.Int64, nil
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("is_current").ColumnType(schema.ColumnTypeString).Description("\"Y\" if is the current role of authenticated user, otherwise \"F\".").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {
				return result.(Role).IsCurrent.String, nil
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).Description("The Snowflake region in which the account is located.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {
				return taskClient.(*snowflake_client.Client).Config.Region, nil
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("created_on").ColumnType(schema.ColumnTypeTimestamp).Description("Date and time when the role was created.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {
				return result.(Role).CreatedOn.Time, nil
			})).Build(),
	}
}

func (x *TableSnowflakeRoleGenerator) GetSubTables() []*schema.Table {
	return []*schema.Table{
		table_schema_generator.GenTableSchema(&TableSnowflakeRoleGrantGenerator{}),
	}
}
