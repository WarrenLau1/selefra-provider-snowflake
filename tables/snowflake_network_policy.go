package tables

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/selefra/selefra-provider-snowflake/snowflake_client"

	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
	"github.com/selefra/selefra-provider-snowflake/table_schema_generator"
)

type TableSnowflakeNetworkPolicyGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableSnowflakeNetworkPolicyGenerator{}

func (x *TableSnowflakeNetworkPolicyGenerator) GetTableName() string {
	return "snowflake_network_policy"
}

func (x *TableSnowflakeNetworkPolicyGenerator) GetTableDescription() string {
	return ""
}

func (x *TableSnowflakeNetworkPolicyGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableSnowflakeNetworkPolicyGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableSnowflakeNetworkPolicyGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {

			db, err := snowflake_client.Connect(ctx, taskClient.(*snowflake_client.Client).Config)
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}
			rows, err := db.QueryContext(ctx, "SHOW NETWORK POLICIES")
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}
			defer rows.Close()

			for rows.Next() {
				var Name sql.NullString
				var CreatedOn sql.NullTime
				var Comment sql.NullString
				var EntriesInAllowedIPList sql.NullInt64
				var EntriesInBlockedIPList sql.NullInt64

				err = rows.Scan(&CreatedOn, &Name, &Comment, &EntriesInAllowedIPList, &EntriesInBlockedIPList)
				if err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)
				}
				resultChannel <- NetworkPolicy{Name, CreatedOn, Comment, EntriesInAllowedIPList, EntriesInBlockedIPList}
			}

			for rows.NextResultSet() {
				for rows.Next() {
					var Name sql.NullString
					var CreatedOn sql.NullTime
					var Comment sql.NullString
					var EntriesInAllowedIPList sql.NullInt64
					var EntriesInBlockedIPList sql.NullInt64

					err = rows.Scan(&CreatedOn, &Name, &Comment, &EntriesInAllowedIPList, &EntriesInBlockedIPList)
					if err != nil {
						return schema.NewDiagnosticsErrorPullTable(task.Table, err)
					}
					resultChannel <- NetworkPolicy{Name, CreatedOn, Comment, EntriesInAllowedIPList, EntriesInBlockedIPList}
				}
			}

			return schema.NewDiagnosticsErrorPullTable(task.Table, nil)

		},
	}
}

type NetworkPolicy struct {
	Name                   sql.NullString `json:"name"`
	CreatedOn              sql.NullTime   `json:"created_on"`
	Comment                sql.NullString `json:"comment"`
	EntriesInAllowedIPList sql.NullInt64  `json:"entries_in_allowed_ip_list"`
	EntriesInBlockedIPList sql.NullInt64  `json:"entries_in_blocked_ip_list"`
}

func DescribeNetworkPolicy(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, error) {
	var policyName string
	if result != nil {
		policyName = result.(NetworkPolicy).Name.String
	} else {
		// todo
		//policyName = d.KeyColumnQualString("name")
	}

	if policyName == "" {
		return nil, nil
	}

	db, err := snowflake_client.Connect(ctx, taskClient.(*snowflake_client.Client).Config)
	if err != nil {

		return nil, err
	}
	rows, err := db.QueryContext(ctx, fmt.Sprintf("DESCRIBE NETWORK POLICY %s", policyName))
	if err != nil {

		return nil, err
	}
	defer rows.Close()

	networkIPlist := map[string]string{}
	for rows.Next() {
		var name sql.NullString
		var value sql.NullString

		err = rows.Scan(&name, &value)
		if err != nil {

			return nil, err
		}
		networkIPlist[name.String] = value.String
	}
	return networkIPlist, nil
}

func (x *TableSnowflakeNetworkPolicyGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return nil
}

func (x *TableSnowflakeNetworkPolicyGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).Description("Identifier for the network policy.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {
				return result.(NetworkPolicy).Name.String, nil
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("comment").ColumnType(schema.ColumnTypeString).Description("Specifies a comment for the network policy.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {
				return result.(NetworkPolicy).Comment.String, nil
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("entries_in_blocked_ip_list").ColumnType(schema.ColumnTypeInt).Description("No of entries in the blocked IP list.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {
				return result.(NetworkPolicy).EntriesInBlockedIPList.Int64, nil
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("allowed_ip_list").ColumnType(schema.ColumnTypeString).Description("Comma-separated list of one or more IPv4 addresses that are allowed access to your Snowflake account.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {
				// 001
				r, err := DescribeNetworkPolicy(ctx, clientMeta, taskClient, task, row, column, result)
				if err != nil {
					return nil, schema.NewDiagnosticsErrorColumnValueExtractor(task.Table, column, err)
				}
				extractor := column_value_extractor.StructSelector("ALLOWED_IP_LIST")
				return extractor.Extract(ctx, clientMeta, taskClient, task, row, column, r)
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("blocked_ip_list").ColumnType(schema.ColumnTypeString).Description("Comma-separated list of one or more IPv4 addresses that are denied access to your Snowflake account.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {
				// 001
				r, err := DescribeNetworkPolicy(ctx, clientMeta, taskClient, task, row, column, result)
				if err != nil {
					return nil, schema.NewDiagnosticsErrorColumnValueExtractor(task.Table, column, err)
				}
				extractor := column_value_extractor.StructSelector("BLOCKED_IP_LIST")
				return extractor.Extract(ctx, clientMeta, taskClient, task, row, column, r)
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).Description("The Snowflake region in which the account is located.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {
				return taskClient.(*snowflake_client.Client).Config.Region, nil
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("account").ColumnType(schema.ColumnTypeString).Description("The Snowflake account ID.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {
				return taskClient.(*snowflake_client.Client).Config.Account, nil
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("created_on").ColumnType(schema.ColumnTypeTimestamp).Description("Date and time when the policy was created.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {
				return result.(NetworkPolicy).CreatedOn.Time, nil
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("entries_in_allowed_ip_list").ColumnType(schema.ColumnTypeInt).Description("No of entries in the allowed IP list.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {
				return result.(NetworkPolicy).EntriesInAllowedIPList.Int64, nil
			})).Build(),
	}
}

func (x *TableSnowflakeNetworkPolicyGenerator) GetSubTables() []*schema.Table {
	return nil
}
