package provider

import (
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-snowflake/table_schema_generator"
	"github.com/selefra/selefra-provider-snowflake/tables"
)

func GenTables() []*schema.Table {
	return []*schema.Table{
		table_schema_generator.GenTableSchema(&tables.TableSnowflakeWarehouseGenerator{}),
		table_schema_generator.GenTableSchema(&tables.TableSnowflakeAccountParameterGenerator{}),
		table_schema_generator.GenTableSchema(&tables.TableSnowflakeResourceMonitorGenerator{}),
		table_schema_generator.GenTableSchema(&tables.TableSnowflakeLoginHistoryGenerator{}),
		table_schema_generator.GenTableSchema(&tables.TableSnowflakeSchemataGenerator{}),
		table_schema_generator.GenTableSchema(&tables.TableSnowflakeSessionPolicyGenerator{}),
		table_schema_generator.GenTableSchema(&tables.TableSnowflakeAccountGrantGenerator{}),
		table_schema_generator.GenTableSchema(&tables.TableSnowflakeUserGenerator{}),
		table_schema_generator.GenTableSchema(&tables.TableSnowflakeRoleGenerator{}),
		table_schema_generator.GenTableSchema(&tables.TableSnowflakeNetworkPolicyGenerator{}),
		table_schema_generator.GenTableSchema(&tables.TableSnowflakeSessionGenerator{}),
		table_schema_generator.GenTableSchema(&tables.TableSnowflakeViewGenerator{}),
		table_schema_generator.GenTableSchema(&tables.TableSnowflakeDatabaseGenerator{}),
	}
}
