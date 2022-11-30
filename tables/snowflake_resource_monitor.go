package tables

import (
	"context"
	"database/sql"
	"github.com/selefra/selefra-provider-snowflake/snowflake_client"

	"github.com/jmoiron/sqlx"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
	"github.com/selefra/selefra-provider-snowflake/table_schema_generator"
)

type TableSnowflakeResourceMonitorGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableSnowflakeResourceMonitorGenerator{}

func (x *TableSnowflakeResourceMonitorGenerator) GetTableName() string {
	return "snowflake_resource_monitor"
}

func (x *TableSnowflakeResourceMonitorGenerator) GetTableDescription() string {
	return ""
}

func (x *TableSnowflakeResourceMonitorGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableSnowflakeResourceMonitorGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableSnowflakeResourceMonitorGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {

			db, err := snowflake_client.Connect(ctx, taskClient.(*snowflake_client.Client).Config)
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}

			rows, err := db.QueryContext(ctx, "SHOW RESOURCE MONITORS")
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}
			defer rows.Close()

			resourceMonitors := []ResourceMonitor{}

			err = sqlx.StructScan(rows, &resourceMonitors)
			if err != nil {
				if err == sql.ErrNoRows {
					return schema.NewDiagnosticsErrorPullTable(task.Table, nil)
				}
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}

			for _, resourceMonitor := range resourceMonitors {
				resultChannel <- resourceMonitor
			}
			return schema.NewDiagnosticsErrorPullTable(task.Table, nil)

		},
	}
}

func (x *TableSnowflakeResourceMonitorGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return nil
}

func (x *TableSnowflakeResourceMonitorGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("end_time").ColumnType(schema.ColumnTypeTimestamp).Description("Date and time when the monitor was stopped.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("suspend_immediately_at").ColumnType(schema.ColumnTypeJSON).Description("Levels to which to suspend warehouse.").
			Extractor(column_value_extractor.StructSelector("SuspendImmediatelyAt.String")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("suspend_at").ColumnType(schema.ColumnTypeJSON).Description("Levels to which to suspend warehouse.").
			Extractor(column_value_extractor.StructSelector("SuspendAt.String")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("owner").ColumnType(schema.ColumnTypeString).Description("Role that owns the warehouse.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).Description("Name for warehouse.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("credit_quota").ColumnType(schema.ColumnTypeFloat).Description("Specifies the number of Snowflake credits allocated to the monitor for the specified frequency interval.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("used_credits").ColumnType(schema.ColumnTypeFloat).Description("Number of credits used in the current monthly billing cycle by all the warehouses associated with the resource monitor.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("frequency").ColumnType(schema.ColumnTypeString).Description("The interval at which the used credits reset relative to the specified start date (Daily,Weekly,Monthly,Yearly,Never).").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("start_time").ColumnType(schema.ColumnTypeTimestamp).Description("Date and time when the monitor was started.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("notify_at").ColumnType(schema.ColumnTypeJSON).Description("Levels to which to alert.").
			Extractor(column_value_extractor.StructSelector("NotifyAt.String")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("notify_users").ColumnType(schema.ColumnTypeString).Description("Who to notify when alerting.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).Description("The Snowflake region in which the account is located.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {
				return taskClient.(*snowflake_client.Client).Config.Region, nil
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("remaining_credits").ColumnType(schema.ColumnTypeFloat).Description("Number of credits still available to use in the current monthly billing cycle.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("level").ColumnType(schema.ColumnTypeString).Description("Specifies whether the resource monitor is used to monitor the credit usage for your entire Account (i.e. all warehouses in the account) or a specific set of individual warehouses.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("created_on").ColumnType(schema.ColumnTypeTimestamp).Description("Date and time when the monitor was created.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("comment").ColumnType(schema.ColumnTypeString).Description("Comment for the warehouse.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("account").ColumnType(schema.ColumnTypeString).Description("The Snowflake account ID.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {
				return taskClient.(*snowflake_client.Client).Config.Account, nil
			})).Build(),
	}
}

func (x *TableSnowflakeResourceMonitorGenerator) GetSubTables() []*schema.Table {
	return nil
}
