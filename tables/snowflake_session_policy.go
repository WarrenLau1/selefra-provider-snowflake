package tables

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/selefra/selefra-provider-snowflake/snowflake_client"

	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
	"github.com/selefra/selefra-provider-snowflake/table_schema_generator"
	"github.com/snowflakedb/gosnowflake"
)

type TableSnowflakeSessionPolicyGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableSnowflakeSessionPolicyGenerator{}

func (x *TableSnowflakeSessionPolicyGenerator) GetTableName() string {
	return "snowflake_session_policy"
}

func (x *TableSnowflakeSessionPolicyGenerator) GetTableDescription() string {
	return ""
}

func (x *TableSnowflakeSessionPolicyGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableSnowflakeSessionPolicyGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableSnowflakeSessionPolicyGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {

			db, err := snowflake_client.Connect(ctx, taskClient.(*snowflake_client.Client).Config)
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}
			rows, err := db.QueryContext(ctx, "SHOW SESSION POLICIES")
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}
			columns, err := rows.Columns()
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}

			for rows.Next() {
				policy := SessionPolicy{}

				cols := make([]interface{}, len(columns))

				for i, col := range columns {
					cols[i] = SessionPolicyCol(col, &policy)
				}

				err = rows.Scan(cols...)
				if err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)
				}
				resultChannel <- policy
			}

			for rows.NextResultSet() {
				for rows.Next() {
					policy := SessionPolicy{}

					cols := make([]interface{}, len(columns))

					for i, col := range columns {
						cols[i] = SessionPolicyCol(col, &policy)
					}

					err = rows.Scan(cols...)
					if err != nil {
						return schema.NewDiagnosticsErrorPullTable(task.Table, err)
					}
					resultChannel <- policy
				}
			}
			return schema.NewDiagnosticsErrorPullTable(task.Table, nil)

		},
	}
}

type Policy struct {
	CreatedOn    sql.NullString `db:"created_on"`
	Name         sql.NullString `db:"name"`
	DatabaseName sql.NullString `db:"database_name"`
	SchemaName   sql.NullString `db:"schema_name"`
	Kind         sql.NullString `db:"kind"`
	Owner        sql.NullString `db:"owner"`
	Comment      sql.NullString `db:"comment"`
}
type SessionPolicy Policy

func DescribeSessionPolicy(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, error) {
	var policy SessionPolicy
	if result != nil {
		policy = result.(SessionPolicy)
	}

	if !policy.Name.Valid {
		return nil, nil
	}

	db, err := snowflake_client.Connect(ctx, taskClient.(*snowflake_client.Client).Config)
	if err != nil {

		return nil, err
	}
	rows, err := db.QueryContext(ctx, fmt.Sprintf("DESCRIBE SESSION POLICY %s.%s.%s", policy.DatabaseName.String, policy.SchemaName.String, policy.Name.String))
	if err != nil {
		if err.(*gosnowflake.SnowflakeError) != nil {

			return nil, nil
		}
		return nil, err
	}
	defer rows.Close()

	policyProperties := struct {
		SessionIdleTimeoutMins   sql.NullInt64 `json:"session_idle_timeout_mins"`
		SessionUiIdleTimeoutMins sql.NullInt64 `json:"session_ui_idle_timeout_mins"`
	}{}

	for rows.Next() {
		var created_on sql.NullTime
		var name sql.NullString
		var session_idle_timeout_mins sql.NullInt64
		var session_ui_idle_timeout_mins sql.NullInt64
		var comment sql.NullString

		err = rows.Scan(&created_on, &name, &session_idle_timeout_mins, &session_ui_idle_timeout_mins, &comment)
		if err != nil {

			return nil, err
		}
		policyProperties.SessionIdleTimeoutMins = session_idle_timeout_mins
		policyProperties.SessionUiIdleTimeoutMins = session_ui_idle_timeout_mins
	}
	return policyProperties, nil
}

// SessionPolicyCol returns a reference for a column of a SessionPolicy
func SessionPolicyCol(colname string, sp *SessionPolicy) interface{} {
	switch colname {
	case "created_on":
		return &sp.CreatedOn
	case "name":
		return &sp.Name
	case "database_name":
		return &sp.DatabaseName
	case "schema_name":
		return &sp.SchemaName
	case "kind":
		return &sp.Kind
	case "owner":
		return &sp.Owner
	case "comment":
		return &sp.Comment
	default:
		panic("unknown column " + colname)
	}
}

func (x *TableSnowflakeSessionPolicyGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return nil
}

func (x *TableSnowflakeSessionPolicyGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("created_on").ColumnType(schema.ColumnTypeTimestamp).Description("Date and time of the creation of session policy.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {
				return result.(SessionPolicy).CreatedOn.String, nil
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("schema_name").ColumnType(schema.ColumnTypeString).Description("Name of the schema in database policy belongs.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {
				return result.(SessionPolicy).SchemaName.String, nil
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("comment").ColumnType(schema.ColumnTypeString).Description("Comment for this policy.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {
				return result.(SessionPolicy).Comment.String, nil
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).Description("The Snowflake region in which the account is located.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {
				return taskClient.(*snowflake_client.Client).Config.Region, nil
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).Description("Identifier for the session policy.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {
				return result.(SessionPolicy).Name.String, nil
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("database_name").ColumnType(schema.ColumnTypeString).Description("Name of the database policy belongs.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {
				return result.(SessionPolicy).DatabaseName.String, nil
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("kind").ColumnType(schema.ColumnTypeString).Description("Type of the snowflake policy.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {
				return result.(SessionPolicy).Kind.String, nil
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("owner").ColumnType(schema.ColumnTypeString).Description("Name of the role that owns the policy.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {
				return result.(SessionPolicy).Owner.String, nil
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("session_idle_timeout_mins").ColumnType(schema.ColumnTypeInt).Description("Time period in minutes of inactivity with either the web interface or a programmatic client.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {
				// 004
				r, err := DescribeSessionPolicy(ctx, clientMeta, taskClient, task, row, column, result)
				if err != nil {
					return nil, schema.NewDiagnosticsErrorColumnValueExtractor(task.Table, column, err)
				}
				return column_value_extractor.DefaultColumnValueExtractor.Extract(ctx, clientMeta, taskClient, task, row, column, r)
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("session_ui_idle_timeout_mins").ColumnType(schema.ColumnTypeInt).Description("Time period in minutes of inactivity with the web interface.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {
				// 004
				r, err := DescribeSessionPolicy(ctx, clientMeta, taskClient, task, row, column, result)
				if err != nil {
					return nil, schema.NewDiagnosticsErrorColumnValueExtractor(task.Table, column, err)
				}
				return column_value_extractor.DefaultColumnValueExtractor.Extract(ctx, clientMeta, taskClient, task, row, column, r)
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("account").ColumnType(schema.ColumnTypeString).Description("The Snowflake account ID.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {
				return taskClient.(*snowflake_client.Client).Config.Account, nil
			})).Build(),
	}
}

func (x *TableSnowflakeSessionPolicyGenerator) GetSubTables() []*schema.Table {
	return nil
}
