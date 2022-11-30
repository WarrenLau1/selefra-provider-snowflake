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

type TableSnowflakeSessionGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableSnowflakeSessionGenerator{}

func (x *TableSnowflakeSessionGenerator) GetTableName() string {
	return "snowflake_session"
}

func (x *TableSnowflakeSessionGenerator) GetTableDescription() string {
	return ""
}

func (x *TableSnowflakeSessionGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableSnowflakeSessionGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableSnowflakeSessionGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {

			db, err := snowflake_client.Connect(ctx, taskClient.(*snowflake_client.Client).Config)
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}

			condition := ""
			query := "SELECT * FROM SNOWFLAKE.ACCOUNT_USAGE.SESSIONS"
			if condition != "" {
				query = fmt.Sprintf("%s where %s;", query, condition)
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
				session := Session{}

				cols := make([]interface{}, len(columns))
				for i, col := range columns {
					cols[i] = SessionCol(col, &session)
				}

				err = rows.Scan(cols...)
				if err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)
				}
				resultChannel <- session
			}

			for rows.NextResultSet() {
				for rows.Next() {
					session := Session{}

					cols := make([]interface{}, len(columns))
					for i, col := range columns {
						cols[i] = SessionCol(col, &session)
					}

					err = rows.Scan(cols...)
					if err != nil {
						return schema.NewDiagnosticsErrorPullTable(task.Table, err)
					}
					resultChannel <- session
				}
			}
			return schema.NewDiagnosticsErrorPullTable(task.Table, nil)

		},
	}
}

type ResourceMonitor struct {
	Name                 sql.NullString `json:"name" db:"name"`
	CreditQuota          sql.NullString `json:"credit_quota" db:"credit_quota"`
	UsedCredits          sql.NullString `json:"used_credits" db:"used_credits"`
	RemainingCredits     sql.NullString `json:"remaining_credits" db:"remaining_credits"`
	Level                sql.NullString `json:"level" db:"level"`
	Frequency            sql.NullString `json:"frequency" db:"frequency"`
	StartTime            sql.NullTime   `json:"start_time" db:"start_time"`
	EndTime              sql.NullTime   `json:"end_time" db:"end_time"`
	NotifyAt             sql.NullString `json:"notify_at" db:"notify_at"`
	SuspendAt            sql.NullString `json:"suspend_at" db:"suspend_at"`
	SuspendImmediatelyAt sql.NullString `json:"suspend_immediately_at" db:"suspend_immediately_at"`
	CreatedOn            sql.NullTime   `json:"created_on" db:"created_on"`
	Owner                sql.NullString `json:"owner" db:"owner"`
	Comment              sql.NullString `json:"comment" db:"comment"`
	NotifyUsers          sql.NullString `json:"notify_users" db:"notify_users"`
}
type Result struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int    `json:"expires_in"`
}
type Role struct {
	CreatedOn       sql.NullTime   `json:"created_on"`
	Name            sql.NullString `json:"name"`
	IsDefault       sql.NullString `json:"is_default"`
	IsCurrent       sql.NullString `json:"is_current"`
	IsInherited     sql.NullString `json:"is_inherited"`
	AssignedToUsers sql.NullInt64  `json:"assigned_to_users"`
	GrantedToRoles  sql.NullInt64  `json:"granted_to_roles"`
	GrantedRoles    sql.NullInt64  `json:"granted_roles"`
	Owner           sql.NullString `json:"owner"`
	Comment         sql.NullString `json:"comment"`
}
type Session struct {
	SessionId                sql.NullInt64  `json:"SESSION_ID"`
	CreatedOn                sql.NullTime   `json:"CREATED_ON"`
	UserName                 sql.NullString `json:"USER_NAME"`
	AuthenticationMethod     sql.NullString `json:"AUTHENTICATION_METHOD"`
	LoginEventId             sql.NullInt64  `json:"LOGIN_EVENT_ID"`
	ClientApplicationVersion sql.NullString `json:"CLIENT_APPLICATION_VERSION"`
	ClientApplicationId      sql.NullString `json:"CLIENT_APPLICATION_ID"`
	ClientEnvironment        sql.NullString `json:"CLIENT_ENVIRONMENT"`
	ClientBuildId            sql.NullString `json:"CLIENT_BUILD_ID"`
	ClientVersion            sql.NullString `json:"CLIENT_VERSION"`
}
type User struct {
	Name                  sql.NullString `json:"name"`
	CreatedOn             sql.NullTime   `json:"created_on"`
	LoginName             sql.NullString `json:"login_name"`
	DisplayName           sql.NullString `json:"display_name"`
	FirstName             sql.NullString `json:"first_name"`
	LastName              sql.NullString `json:"last_name"`
	Email                 sql.NullString `json:"email"`
	MinsToUnlock          sql.NullString `json:"mins_to_unlock"`
	DaysToExpiry          sql.NullString `json:"days_to_expiry"`
	Comment               sql.NullString `json:"comment"`
	Disabled              sql.NullString `json:"disabled"`
	MustChangePassword    sql.NullString `json:"must_change_password"`
	SnowflakeLock         sql.NullString `json:"snowflake_lock"`
	DefaultWarehouse      sql.NullString `json:"default_warehouse"`
	DefaultNamespace      sql.NullString `json:"default_namespace"`
	DefaultRole           sql.NullString `json:"default_role"`
	DefaultSecondaryRoles sql.NullString `json:"default_secondary_roles"`
	ExtAuthnDuo           sql.NullString `json:"ext_authn_duo"`
	ExtAuthnUid           sql.NullString `json:"ext_authn_uid"`
	MinsToBypassMFA       sql.NullString `json:"mins_to_bypass_mfa"`
	Owner                 sql.NullString `json:"owner"`
	LastSuccessLogin      sql.NullTime   `json:"last_success_login"`
	ExpiresAtTime         sql.NullTime   `json:"expires_at_time"`
	LockedUntilTime       sql.NullTime   `json:"locked_until_time"`
	HasPassword           sql.NullString `json:"has_password"`
	HasRsaPublicKey       sql.NullString `json:"has_rsa_public_key"`
}
type Warehouse struct {
	Name            sql.NullString `json:"name" db:"name"`
	State           sql.NullString `json:"state" db:"state"`
	Type            sql.NullString `json:"type" db:"type"`
	Size            sql.NullString `json:"size" db:"size"`
	MinClusterCount sql.NullInt64  `json:"min_cluster_count" db:"min_cluster_count"`
	MaxClusterCount sql.NullInt64  `json:"max_cluster_count" db:"max_cluster_count"`
	StartedClusters sql.NullInt64  `json:"started_clusters" db:"started_clusters"`
	Running         sql.NullInt64  `json:"running" db:"running"`
	Queued          sql.NullInt64  `json:"queued" db:"queued"`
	IsDefault       sql.NullString `json:"is_default" db:"is_default"`
	IsCurrent       sql.NullString `json:"is_current" db:"is_current"`
	AutoSuspend     sql.NullInt64  `json:"auto_suspend" db:"auto_suspend"`
	AutoResume      sql.NullBool   `json:"auto_resume" db:"auto_resume"`
	Available       sql.NullString `json:"available" db:"available"`
	Provisioning    sql.NullString `json:"provisioning" db:"provisioning"`
	Quiescing       sql.NullString `json:"quiescing" db:"quiescing"`
	Other           sql.NullString `json:"other" db:"other"`
	CreatedOn       sql.NullTime   `json:"created_on" db:"created_on"`
	ResumedOn       sql.NullTime   `json:"resumed_on" db:"resumed_on"`
	UpdatedOn       sql.NullTime   `json:"updated_on" db:"updated_on"`
	Owner           sql.NullString `json:"owner" db:"owner"`
	Comment         sql.NullString `json:"comment" db:"comment"`
	ResourceMonitor sql.NullString `json:"resource_monitor" db:"resource_monitor"`
	Actives         sql.NullInt64  `json:"actives" db:"actives"`
	Pendings        sql.NullInt64  `json:"pendings" db:"pendings"`
	Failed          sql.NullInt64  `json:"failed" db:"failed"`
	Suspended       sql.NullInt64  `json:"suspended" db:"suspended"`
	UUID            sql.NullString `json:"uuid" db:"uuid"`
	ScalingPolicy   sql.NullString `json:"scaling_policy" db:"scaling_policy"`
}

// SessionCol returns a reference for a column of a Session
func SessionCol(colname string, item *Session) interface{} {
	switch colname {
	case "SESSION_ID":
		return &item.SessionId
	case "CREATED_ON":
		return &item.CreatedOn
	case "USER_NAME":
		return &item.UserName
	case "AUTHENTICATION_METHOD":
		return &item.AuthenticationMethod
	case "LOGIN_EVENT_ID":
		return &item.LoginEventId
	case "CLIENT_APPLICATION_VERSION":
		return &item.ClientApplicationVersion
	case "CLIENT_APPLICATION_ID":
		return &item.ClientApplicationId
	case "CLIENT_ENVIRONMENT":
		return &item.ClientEnvironment
	case "CLIENT_BUILD_ID":
		return &item.ClientBuildId
	case "CLIENT_VERSION":
		return &item.ClientVersion
	default:
		panic("unknown column " + colname)
	}
}

func (x *TableSnowflakeSessionGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return nil
}

func (x *TableSnowflakeSessionGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("authentication_method").ColumnType(schema.ColumnTypeString).Description("The authentication method used to access Snowflake.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {
				return result.(Session).AuthenticationMethod.String, nil
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("client_application_id").ColumnType(schema.ColumnTypeString).Description("The identifier for the Snowflake-provided client application used to create the remote session to Snowflake (e.g. JDBC 3.8.7)").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {
				return result.(Session).ClientApplicationId.String, nil
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("client_application_version").ColumnType(schema.ColumnTypeString).Description("The version number (e.g. 3.8.7) of the Snowflake-provided client application used to create the remote session to Snowflake.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {
				return result.(Session).ClientApplicationVersion.String, nil
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).Description("The Snowflake region in which the account is located.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {
				return taskClient.(*snowflake_client.Client).Config.Region, nil
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("account").ColumnType(schema.ColumnTypeString).Description("The Snowflake account ID.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {
				return taskClient.(*snowflake_client.Client).Config.Account, nil
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("client_version").ColumnType(schema.ColumnTypeString).Description("The version number (e.g. 47154) of the third-party client application that uses a Snowflake-provided client to create a remote session to Snowflake, if available.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {
				return result.(Session).ClientVersion.String, nil
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("login_event_id").ColumnType(schema.ColumnTypeString).Description("The unique identifier for the login event.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {
				return strconv.FormatInt(result.(Session).LoginEventId.Int64, 10), nil
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("session_id").ColumnType(schema.ColumnTypeString).Description("The unique identifier for the current session.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {
				return strconv.FormatInt(result.(Session).SessionId.Int64, 10), nil
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("user_name").ColumnType(schema.ColumnTypeString).Description("The user name of the user.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {
				return result.(Session).UserName.String, nil
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("created_on").ColumnType(schema.ColumnTypeTimestamp).Description("Date and time when the session was created.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {
				return result.(Session).CreatedOn.Time, nil
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("client_build_id").ColumnType(schema.ColumnTypeString).Description("The build number (e.g. 41897) of the third-party client application used to create a remote session to Snowflake, if available. For example, a third-party Java application that uses the JDBC driver to connect to Snowflake.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {
				return result.(Session).ClientBuildId.String, nil
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("client_environment").ColumnType(schema.ColumnTypeJSON).Description("The environment variables (e.g. operating system, OCSP mode) of the client used to create a remote session to Snowflake.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {
				return result.(Session).ClientEnvironment.String, nil
			})).Build(),
	}
}

func (x *TableSnowflakeSessionGenerator) GetSubTables() []*schema.Table {
	return nil
}
