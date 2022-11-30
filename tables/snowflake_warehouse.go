package tables

import (
	"context"
	"database/sql"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
	"github.com/selefra/selefra-provider-snowflake/snowflake_client"

	"github.com/jmoiron/sqlx"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-snowflake/table_schema_generator"
)

type TableSnowflakeWarehouseGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableSnowflakeWarehouseGenerator{}

func (x *TableSnowflakeWarehouseGenerator) GetTableName() string {
	return "snowflake_warehouse"
}

func (x *TableSnowflakeWarehouseGenerator) GetTableDescription() string {
	return ""
}

func (x *TableSnowflakeWarehouseGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableSnowflakeWarehouseGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableSnowflakeWarehouseGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {

			db, err := snowflake_client.Connect(ctx, taskClient.(*snowflake_client.Client).Config)
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}

			rows, err := db.QueryContext(ctx, "SHOW WAREHOUSES")
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}
			defer rows.Close()

			dbs := []Warehouse{}

			err = sqlx.StructScan(rows, &dbs)
			if err != nil {
				if err == sql.ErrNoRows {
					return schema.NewDiagnosticsErrorPullTable(task.Table, nil)
				}
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}

			for _, warehouse := range dbs {
				resultChannel <- warehouse
			}
			return schema.NewDiagnosticsErrorPullTable(task.Table, nil)

		},
	}
}

func (x *TableSnowflakeWarehouseGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return nil
}

func (x *TableSnowflakeWarehouseGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("max_cluster_count").ColumnType(schema.ColumnTypeInt).Description("Maximum number of warehouses for the (multi-cluster) warehouse (always 1 for single warehouses).").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {
				return result.(Warehouse).MaxClusterCount.Int64, nil
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("running").ColumnType(schema.ColumnTypeInt).Description("Number of SQL statements that are being executed by the warehouse.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {
				return result.(Warehouse).Running.Int64, nil
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("created_on").ColumnType(schema.ColumnTypeTimestamp).Description("Date and time when the warehouse was created.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {
				return result.(Warehouse).CreatedOn.Time, nil
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("resumed_on").ColumnType(schema.ColumnTypeTimestamp).Description("Date and time when the warehouse was last started or restarted.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {
				return result.(Warehouse).ResumedOn.Time, nil
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("resource_monitor").ColumnType(schema.ColumnTypeString).Description("ID of resource monitor explicitly assigned to the warehouse; controls the monthly credit usage for the warehouse.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {
				return result.(Warehouse).ResourceMonitor.String, nil
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("quiescing").ColumnType(schema.ColumnTypeString).Description("Percentage of the warehouse compute resources that are executing SQL statements, but will be shut down once the queries complete.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {
				return result.(Warehouse).Quiescing.String, nil
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("type").ColumnType(schema.ColumnTypeString).Description("Warehouse type; STANDARD is the only currently supported type.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {
				return result.(Warehouse).Type.String, nil
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("started_clusters").ColumnType(schema.ColumnTypeInt).Description("Number of warehouses currently started.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {
				return result.(Warehouse).StartedClusters.Int64, nil
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("is_default").ColumnType(schema.ColumnTypeString).Description("Whether the warehouse is the default for the current user.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {
				return result.(Warehouse).IsDefault.String, nil
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("auto_suspend").ColumnType(schema.ColumnTypeInt).Description("Specifies the number of seconds of inactivity after which a warehouse is automatically suspended.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {
				return result.(Warehouse).AutoSuspend.Int64, nil
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("auto_resume").ColumnType(schema.ColumnTypeBool).Description("Specifies whether to automatically resume a warehouse when a SQL statement (e.g. query) is submitted to it.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {
				return result.(Warehouse).AutoResume.Bool, nil
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("available").ColumnType(schema.ColumnTypeString).Description("Percentage of the warehouse compute resources that are provisioned and available.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {
				return result.(Warehouse).Available.String, nil
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("provisioning").ColumnType(schema.ColumnTypeString).Description("Percentage of the warehouse compute resources that are in the process of provisioning.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {
				return result.(Warehouse).Provisioning.String, nil
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("updated_on").ColumnType(schema.ColumnTypeTimestamp).Description("Date and time when the warehouse was last updated, which includes changing any of the properties of the warehouse or changing the state (STARTED, SUSPENDED, RESIZING) of the warehouse.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {
				return result.(Warehouse).UpdatedOn.Time, nil
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("owner").ColumnType(schema.ColumnTypeString).Description("Role that owns the warehouse.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {
				return result.(Warehouse).Owner.String, nil
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("scaling_policy").ColumnType(schema.ColumnTypeString).Description("Policy that determines when additional warehouses (in a multi-cluster warehouse) are automatically started and shut down.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {
				return result.(Warehouse).ScalingPolicy.String, nil
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("min_cluster_count").ColumnType(schema.ColumnTypeInt).Description("Minimum number of warehouses for the (multi-cluster) warehouse (always 1 for single warehouses).").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {
				return result.(Warehouse).MinClusterCount.Int64, nil
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("queued").ColumnType(schema.ColumnTypeInt).Description("Number of SQL statements that are queued for the warehouse.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {
				return result.(Warehouse).Queued.Int64, nil
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("other").ColumnType(schema.ColumnTypeString).Description("Percentage of the warehouse compute resources that are in a state other than available, provisioning, or quiescing.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {
				return result.(Warehouse).Other.String, nil
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).Description("The Snowflake region in which the account is located.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {
				return taskClient.(*snowflake_client.Client).Config.Region, nil
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).Description("Name for warehouse.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {
				return result.(Warehouse).Name.String, nil
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("state").ColumnType(schema.ColumnTypeString).Description("Whether the warehouse is active/running (STARTED), inactive (SUSPENDED), or resizing (RESIZING).").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {
				return result.(Warehouse).State.String, nil
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("size").ColumnType(schema.ColumnTypeString).Description("Size of the warehouse (X-Small, Small, Medium, Large, X-Large, etc.)").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {
				return result.(Warehouse).Size.String, nil
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("is_current").ColumnType(schema.ColumnTypeString).Description("Whether the warehouse is in use for the session.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {
				return result.(Warehouse).IsCurrent.String, nil
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("comment").ColumnType(schema.ColumnTypeString).Description("Comment for the warehouse.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {
				return result.(Warehouse).Comment.String, nil
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("account").ColumnType(schema.ColumnTypeString).Description("The Snowflake account ID.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {
				return taskClient.(*snowflake_client.Client).Config.Account, nil
			})).Build(),
	}
}

func (x *TableSnowflakeWarehouseGenerator) GetSubTables() []*schema.Table {
	return nil
}
