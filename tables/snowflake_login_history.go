package tables

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
	"github.com/selefra/selefra-provider-snowflake/snowflake_client"
	"github.com/selefra/selefra-provider-snowflake/table_schema_generator"
	"strconv"
)

type TableSnowflakeLoginHistoryGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableSnowflakeLoginHistoryGenerator{}

func (x *TableSnowflakeLoginHistoryGenerator) GetTableName() string {
	return "snowflake_login_history"
}

func (x *TableSnowflakeLoginHistoryGenerator) GetTableDescription() string {
	return ""
}

func (x *TableSnowflakeLoginHistoryGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableSnowflakeLoginHistoryGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableSnowflakeLoginHistoryGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {

			db, err := snowflake_client.Connect(ctx, taskClient.(*snowflake_client.Client).Config)
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}

			condition := ""
			query := "SELECT * FROM SNOWFLAKE.ACCOUNT_USAGE.LOGIN_HISTORY"

			if condition != "" {
				query = fmt.Sprintf("%s where %s", query, condition)
			}

			rows, err := db.QueryContext(ctx, query)
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}
			defer rows.Close()

			columns, err := rows.Columns()
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}

			for rows.Next() {
				loginHistory := LoginHistory{}

				cols := make([]interface{}, len(columns))
				for i, col := range columns {
					cols[i] = LoginHistoryCol(col, &loginHistory)
				}

				err = rows.Scan(cols...)
				if err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)
				}
				resultChannel <- loginHistory
			}

			for rows.NextResultSet() {
				for rows.Next() {
					loginHistory := LoginHistory{}

					cols := make([]interface{}, len(columns))
					for i, col := range columns {
						cols[i] = LoginHistoryCol(col, &loginHistory)
					}

					err = rows.Scan(cols...)
					if err != nil {
						return schema.NewDiagnosticsErrorPullTable(task.Table, err)
					}
					resultChannel <- loginHistory
				}
			}
			return schema.NewDiagnosticsErrorPullTable(task.Table, nil)

		},
	}
}

type LoginHistory struct {
	EventId                    sql.NullInt64  `json:"EVENT_ID"`
	EventTimestamp             sql.NullTime   `json:"EVENT_TIMESTAMP"`
	EventType                  sql.NullString `json:"EVENT_TYPE"`
	UserName                   sql.NullString `json:"USER_NAME"`
	ClientIp                   sql.NullString `json:"CLIENT_IP"`
	ReportedClientType         sql.NullString `json:"REPORTED_CLIENT_TYPE"`
	ReportedClientVersion      sql.NullString `json:"REPORTED_CLIENT_VERSION"`
	FirstAuthenticationFactor  sql.NullString `json:"FIRST_AUTHENTICATION_FACTOR"`
	SecondAuthenticationFactor sql.NullString `json:"SECOND_AUTHENTICATION_FACTOR"`
	IsSuccess                  sql.NullString `json:"IS_SUCCESS"`
	ErrorCode                  sql.NullInt64  `json:"ERROR_CODE"`
	ErrorMessage               sql.NullString `json:"ERROR_MESSAGE"`
	RelatedEventId             sql.NullInt64  `json:"RELATED_EVENT_ID"`
	Connection                 sql.NullInt64  `json:"CONNECTION"`
}

// LoginHistoryCol returns a reference for a column of a LoginHistory
func LoginHistoryCol(colname string, item *LoginHistory) interface{} {
	switch colname {
	case "EVENT_ID":
		return &item.EventId
	case "EVENT_TIMESTAMP":
		return &item.EventTimestamp
	case "EVENT_TYPE":
		return &item.EventType
	case "USER_NAME":
		return &item.UserName
	case "CLIENT_IP":
		return &item.ClientIp
	case "REPORTED_CLIENT_TYPE":
		return &item.ReportedClientType
	case "REPORTED_CLIENT_VERSION":
		return &item.ReportedClientVersion
	case "FIRST_AUTHENTICATION_FACTOR":
		return &item.FirstAuthenticationFactor
	case "SECOND_AUTHENTICATION_FACTOR":
		return &item.SecondAuthenticationFactor
	case "IS_SUCCESS":
		return &item.IsSuccess
	case "ERROR_CODE":
		return &item.ErrorCode
	case "ERROR_MESSAGE":
		return &item.ErrorMessage
	case "RELATED_EVENT_ID":
		return &item.RelatedEventId
	case "CONNECTION":
		return &item.Connection
	default:
		panic("unknown column " + colname)
	}
}

func (x *TableSnowflakeLoginHistoryGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return nil
}

func (x *TableSnowflakeLoginHistoryGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("account").ColumnType(schema.ColumnTypeString).Description("The Snowflake account ID.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {
				return taskClient.(*snowflake_client.Client).Config.Account, nil
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("reported_client_version").ColumnType(schema.ColumnTypeString).Description("Reported version of the client software. This information is not authenticated.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {
				return result.(LoginHistory).ReportedClientVersion.String, nil
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("second_authentication_factor").ColumnType(schema.ColumnTypeString).Description("The second factor, if using multi factor authentication, or NULL otherwise.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {
				return result.(LoginHistory).SecondAuthenticationFactor.String, nil
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("is_success").ColumnType(schema.ColumnTypeString).Description("Whether the user's request was successful or not.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {
				return result.(LoginHistory).IsSuccess.String, nil
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("error_code").ColumnType(schema.ColumnTypeInt).Description("Error code, if the request was not successful.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {
				return result.(LoginHistory).ErrorCode.Int64, nil
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("event_id").ColumnType(schema.ColumnTypeString).Description("Internal/system-generated identifier for the login attempt.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {
				return strconv.FormatInt(result.(LoginHistory).EventId.Int64, 10), nil
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("client_ip").ColumnType(schema.ColumnTypeString).Description("IP address where the request originated from.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {
				return result.(LoginHistory).ClientIp.String, nil
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("first_authentication_factor").ColumnType(schema.ColumnTypeString).Description("Method used to authenticate the user (the first factor, if using multi factor authentication).").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {
				return result.(LoginHistory).FirstAuthenticationFactor.String, nil
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).Description("The Snowflake region in which the account is located.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {
				return taskClient.(*snowflake_client.Client).Config.Region, nil
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("user_name").ColumnType(schema.ColumnTypeString).Description("User associated with this event.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {
				return result.(LoginHistory).UserName.String, nil
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("event_type").ColumnType(schema.ColumnTypeString).Description("Event type, such as LOGIN for authentication events.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {
				return result.(LoginHistory).EventType.String, nil
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("reported_client_type").ColumnType(schema.ColumnTypeString).Description("Reported type of the client software, such as JDBC_DRIVER, ODBC_DRIVER, etc. This information is not authenticated.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {
				return result.(LoginHistory).ReportedClientType.String, nil
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("error_message").ColumnType(schema.ColumnTypeString).Description("Error message returned to the user, if the request was not successful.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {
				return result.(LoginHistory).ErrorMessage.String, nil
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("related_event_id").ColumnType(schema.ColumnTypeInt).Description("Reserved for future use.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {
				return result.(LoginHistory).RelatedEventId.Int64, nil
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("event_timestamp").ColumnType(schema.ColumnTypeTimestamp).Description("Time (in the UTC time zone) of the event occurrence.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {
				return result.(LoginHistory).EventTimestamp.Time, nil
			})).Build(),
	}
}

func (x *TableSnowflakeLoginHistoryGenerator) GetSubTables() []*schema.Table {
	return nil
}
