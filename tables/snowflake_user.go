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

type TableSnowflakeUserGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableSnowflakeUserGenerator{}

func (x *TableSnowflakeUserGenerator) GetTableName() string {
	return "snowflake_user"
}

func (x *TableSnowflakeUserGenerator) GetTableDescription() string {
	return ""
}

func (x *TableSnowflakeUserGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableSnowflakeUserGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableSnowflakeUserGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {

			db, err := snowflake_client.Connect(ctx, taskClient.(*snowflake_client.Client).Config)
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}
			rows, err := db.QueryContext(ctx, "SHOW USERS")
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}
			defer rows.Close()

			for rows.Next() {
				var Name sql.NullString
				var CreatedOn sql.NullTime
				var LoginName sql.NullString
				var DisplayName sql.NullString
				var FirstName sql.NullString
				var LastName sql.NullString
				var Email sql.NullString
				var MinsToUnlock sql.NullString
				var DaysToExpiry sql.NullString
				var Comment sql.NullString
				var Disabled sql.NullString
				var MustChangePassword sql.NullString
				var SnowflakeLock sql.NullString
				var DefaultWarehouse sql.NullString
				var DefaultNamespace sql.NullString
				var DefaultRole sql.NullString
				var DefaultSecondaryRoles sql.NullString
				var ExtAuthnDuo sql.NullString
				var ExtAuthnUid sql.NullString
				var MinsToBypassMFA sql.NullString
				var Owner sql.NullString
				var LastSuccessLogin sql.NullTime
				var ExpiresAtTime sql.NullTime
				var LockedUntilTime sql.NullTime
				var HasPassword sql.NullString
				var HasRSAPublicKey sql.NullString

				err = rows.Scan(&Name, &CreatedOn, &LoginName, &DisplayName, &FirstName, &LastName, &Email, &MinsToUnlock, &DaysToExpiry, &Comment, &Disabled, &MustChangePassword, &SnowflakeLock, &DefaultWarehouse, &DefaultNamespace, &DefaultRole, &DefaultSecondaryRoles, &ExtAuthnDuo, &ExtAuthnUid, &MinsToBypassMFA, &Owner, &LastSuccessLogin, &ExpiresAtTime, &LockedUntilTime, &HasPassword, &HasRSAPublicKey)
				if err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)
				}
				resultChannel <- User{Name, CreatedOn, LoginName, DisplayName, FirstName, LastName, Email, MinsToUnlock, DaysToExpiry, Comment, Disabled, MustChangePassword, SnowflakeLock, DefaultWarehouse, DefaultNamespace, DefaultRole, DefaultSecondaryRoles, ExtAuthnDuo, ExtAuthnUid, MinsToBypassMFA, Owner, LastSuccessLogin, ExpiresAtTime, LockedUntilTime, HasPassword, HasRSAPublicKey}
			}

			for rows.NextResultSet() {
				for rows.Next() {
					var Name sql.NullString
					var CreatedOn sql.NullTime
					var LoginName sql.NullString
					var DisplayName sql.NullString
					var FirstName sql.NullString
					var LastName sql.NullString
					var Email sql.NullString
					var MinsToUnlock sql.NullString
					var DaysToExpiry sql.NullString
					var Comment sql.NullString
					var Disabled sql.NullString
					var MustChangePassword sql.NullString
					var SnowflakeLock sql.NullString
					var DefaultWarehouse sql.NullString
					var DefaultNamespace sql.NullString
					var DefaultRole sql.NullString
					var DefaultSecondaryRoles sql.NullString
					var ExtAuthnDuo sql.NullString
					var ExtAuthnUid sql.NullString
					var MinsToBypassMFA sql.NullString
					var Owner sql.NullString
					var LastSuccessLogin sql.NullTime
					var ExpiresAtTime sql.NullTime
					var LockedUntilTime sql.NullTime
					var HasPassword sql.NullString
					var HasRSAPublicKey sql.NullString

					err = rows.Scan(&Name, &CreatedOn, &LoginName, &DisplayName, &FirstName, &LastName, &Email, &MinsToUnlock, &DaysToExpiry, &Comment, &Disabled, &MustChangePassword, &SnowflakeLock, &DefaultWarehouse, &DefaultNamespace, &DefaultRole, &DefaultSecondaryRoles, &ExtAuthnDuo, &ExtAuthnUid, &MinsToBypassMFA, &Owner, &LastSuccessLogin, &ExpiresAtTime, &LockedUntilTime, &HasPassword, &HasRSAPublicKey)
					if err != nil {
						return schema.NewDiagnosticsErrorPullTable(task.Table, err)
					}
					resultChannel <- User{Name, CreatedOn, LoginName, DisplayName, FirstName, LastName, Email, MinsToUnlock, DaysToExpiry, Comment, Disabled, MustChangePassword, SnowflakeLock, DefaultWarehouse, DefaultNamespace, DefaultRole, DefaultSecondaryRoles, ExtAuthnDuo, ExtAuthnUid, MinsToBypassMFA, Owner, LastSuccessLogin, ExpiresAtTime, LockedUntilTime, HasPassword, HasRSAPublicKey}
				}
			}
			return schema.NewDiagnosticsErrorPullTable(task.Table, nil)

		},
	}
}

func DescribeUser(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, error) {
	var userName string
	if result != nil {
		userName = result.(User).Name.String
	} else {
		// todo
		//userName = d.KeyColumnQualString("name")
	}

	if userName == "" {
		return nil, nil
	}

	db, err := snowflake_client.Connect(ctx, taskClient.(*snowflake_client.Client).Config)
	if err != nil {

		return nil, err
	}
	rows, err := db.QueryContext(ctx, fmt.Sprintf("DESCRIBE USER %s", userName))
	if err != nil {
		if err.(*gosnowflake.SnowflakeError) != nil {

			return nil, nil
		}

		return nil, err
	}
	defer rows.Close()

	userProperties := map[string]string{}
	for rows.Next() {
		var property sql.NullString
		var value sql.NullString
		var defaultval sql.NullString
		var description sql.NullString

		err = rows.Scan(&property, &value, &defaultval, &description)
		if err != nil {

			return nil, err
		}
		userProperties[property.String] = value.String
	}
	return userProperties, nil
}

func (x *TableSnowflakeUserGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return nil
}

func (x *TableSnowflakeUserGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).Description("Name of the snowflake user.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {
				return result.(User).Name.String, nil
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("comment").ColumnType(schema.ColumnTypeString).Description("Comment associated to user in the dictionary.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {
				return result.(User).Comment.String, nil
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).Description("The Snowflake region in which the account is located.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {
				return taskClient.(*snowflake_client.Client).Config.Region, nil
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("custom_landing_page_url_flush_next_ui_load").ColumnType(schema.ColumnTypeBool).Description("The timestamp on which the last non-null password was set for the user. Default to null if no password has been set yet.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {
				// 001
				r, err := DescribeUser(ctx, clientMeta, taskClient, task, row, column, result)
				if err != nil {
					return nil, schema.NewDiagnosticsErrorColumnValueExtractor(task.Table, column, err)
				}
				extractor := column_value_extractor.StructSelector("CUSTOM_LANDING_PAGE_URL_FLUSH_NEXT_UI_LOAD")
				return extractor.Extract(ctx, clientMeta, taskClient, task, row, column, r)
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("default_role").ColumnType(schema.ColumnTypeString).Description("Primary principal of user session will be set to this role.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {
				return result.(User).DefaultRole.String, nil
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("display_name").ColumnType(schema.ColumnTypeString).Description("Display name of the user.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {
				return result.(User).DisplayName.String, nil
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("first_name").ColumnType(schema.ColumnTypeString).Description("First name of the user.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {
				return result.(User).FirstName.String, nil
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("last_name").ColumnType(schema.ColumnTypeString).Description("Last name of the user.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {
				return result.(User).LastName.String, nil
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("default_namespace").ColumnType(schema.ColumnTypeString).Description("Default database namespace prefix for this user.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {
				return result.(User).DefaultNamespace.String, nil
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("default_secondary_roles").ColumnType(schema.ColumnTypeString).Description("The secondary roles will be set to all roles provided here.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {
				return result.(User).DefaultSecondaryRoles.String, nil
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("mins_to_unlock").ColumnType(schema.ColumnTypeString).Description("Temporary lock on the user will be removed after specified number of minutes.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {
				return result.(User).MinsToUnlock.String, nil
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("snowflake_support").ColumnType(schema.ColumnTypeString).Description("Snowflake Support is allowed to use the user or account.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {
				// 001
				r, err := DescribeUser(ctx, clientMeta, taskClient, task, row, column, result)
				if err != nil {
					return nil, schema.NewDiagnosticsErrorColumnValueExtractor(task.Table, column, err)
				}
				extractor := column_value_extractor.StructSelector("SNOWFLAKE_SUPPORT")
				return extractor.Extract(ctx, clientMeta, taskClient, task, row, column, r)
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("has_password").ColumnType(schema.ColumnTypeBool).Description("Whether the user has password.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {
				return result.(User).HasPassword.String, nil
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("disabled").ColumnType(schema.ColumnTypeString).Description("Whether the user is disabled.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {
				return result.(User).Disabled.String, nil
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("rsa_public_key_2_fp").ColumnType(schema.ColumnTypeString).Description("Fingerprint of user's second RSA public key.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {
				// 001
				r, err := DescribeUser(ctx, clientMeta, taskClient, task, row, column, result)
				if err != nil {
					return nil, schema.NewDiagnosticsErrorColumnValueExtractor(task.Table, column, err)
				}
				extractor := column_value_extractor.StructSelector("RSA_PUBLIC_KEY_2_FP")
				return extractor.Extract(ctx, clientMeta, taskClient, task, row, column, r)
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("login_name").ColumnType(schema.ColumnTypeString).Description("Login name of the user.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {
				return result.(User).LoginName.String, nil
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("created_on").ColumnType(schema.ColumnTypeTimestamp).Description("Timestamp when the user was created.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {
				return result.(User).CreatedOn.Time, nil
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("custom_landing_page_url").ColumnType(schema.ColumnTypeString).Description("Snowflake Support is allowed to use the user or account.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {
				// 001
				r, err := DescribeUser(ctx, clientMeta, taskClient, task, row, column, result)
				if err != nil {
					return nil, schema.NewDiagnosticsErrorColumnValueExtractor(task.Table, column, err)
				}
				extractor := column_value_extractor.StructSelector("CUSTOM_LANDING_PAGE_URL")
				return extractor.Extract(ctx, clientMeta, taskClient, task, row, column, r)
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("expires_at_time").ColumnType(schema.ColumnTypeTimestamp).Description("The date and time when the user's status is set to EXPIRED and the user can no longer log in.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {
				return result.(User).ExpiresAtTime.Time, nil
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("locked_until_time").ColumnType(schema.ColumnTypeTimestamp).Description("Specifies the number of minutes until the temporary lock on the user login is cleared.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {
				return result.(User).LockedUntilTime.Time, nil
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("rsa_public_key_fp").ColumnType(schema.ColumnTypeString).Description("Fingerprint of user's RSA public key.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {
				// 001
				r, err := DescribeUser(ctx, clientMeta, taskClient, task, row, column, result)
				if err != nil {
					return nil, schema.NewDiagnosticsErrorColumnValueExtractor(task.Table, column, err)
				}
				extractor := column_value_extractor.StructSelector("RSA_PUBLIC_KEY_FP")
				return extractor.Extract(ctx, clientMeta, taskClient, task, row, column, r)
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("rsa_public_key_2").ColumnType(schema.ColumnTypeString).Description("Second RSA public key of the user.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {
				// 001
				r, err := DescribeUser(ctx, clientMeta, taskClient, task, row, column, result)
				if err != nil {
					return nil, schema.NewDiagnosticsErrorColumnValueExtractor(task.Table, column, err)
				}
				extractor := column_value_extractor.StructSelector("RSA_PUBLIC_KEY_2")
				return extractor.Extract(ctx, clientMeta, taskClient, task, row, column, r)
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("has_rsa_public_key").ColumnType(schema.ColumnTypeBool).Description("Whether the user has RSA public key.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {
				return result.(User).HasRsaPublicKey.String, nil
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("days_to_expiry").ColumnType(schema.ColumnTypeString).Description("User record will be treated as expired after specified number of days.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {
				return result.(User).DaysToExpiry.String, nil
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("ext_authn_duo").ColumnType(schema.ColumnTypeBool).Description("Whether Duo Security is enabled as second factor authentication.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {
				return result.(User).ExtAuthnDuo.String, nil
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("mins_to_bypass_network_policy").ColumnType(schema.ColumnTypeString).Description("Temporary bypass network policy on the user for a specified number of minutes.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {
				// 001
				r, err := DescribeUser(ctx, clientMeta, taskClient, task, row, column, result)
				if err != nil {
					return nil, schema.NewDiagnosticsErrorColumnValueExtractor(task.Table, column, err)
				}
				extractor := column_value_extractor.StructSelector("MINS_TO_BYPASS_NETWORK_POLICY")
				return extractor.Extract(ctx, clientMeta, taskClient, task, row, column, r)
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("must_change_password").ColumnType(schema.ColumnTypeString).Description("User must change the password.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {
				return result.(User).MustChangePassword.String, nil
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("email").ColumnType(schema.ColumnTypeString).Description("Email address of the user").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {
				return result.(User).Email.String, nil
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("ext_authn_uid").ColumnType(schema.ColumnTypeString).Description("External authentication ID of the user.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {
				return result.(User).ExtAuthnUid.String, nil
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("mins_to_bypass_mfa").ColumnType(schema.ColumnTypeString).Description("Temporary bypass MFA for the user for a specified number of minutes.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {
				return result.(User).MinsToBypassMFA.String, nil
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("snowflake_lock").ColumnType(schema.ColumnTypeString).Description("Whether the user or account is locked by Snowflake.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {
				return result.(User).SnowflakeLock.String, nil
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("owner").ColumnType(schema.ColumnTypeString).Description("Owner of the user in Snowflake.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {
				return result.(User).Owner.String, nil
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("default_warehouse").ColumnType(schema.ColumnTypeString).Description("Default warehouse for this user.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {
				return result.(User).DefaultWarehouse.String, nil
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("last_success_login").ColumnType(schema.ColumnTypeTimestamp).Description("Date and time when the user last logged in to the Snowflake.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {
				return result.(User).LastSuccessLogin.Time, nil
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("password_last_set_time").ColumnType(schema.ColumnTypeString).Description("The timestamp on which the last non-null password was set for the user. Default to null if no password has been set yet.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {
				// 001
				r, err := DescribeUser(ctx, clientMeta, taskClient, task, row, column, result)
				if err != nil {
					return nil, schema.NewDiagnosticsErrorColumnValueExtractor(task.Table, column, err)
				}
				extractor := column_value_extractor.StructSelector("PASSWORD_LAST_SET_TIME")
				return extractor.Extract(ctx, clientMeta, taskClient, task, row, column, r)
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("rsa_public_key").ColumnType(schema.ColumnTypeString).Description("RSA public key of the user.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {
				// 001
				r, err := DescribeUser(ctx, clientMeta, taskClient, task, row, column, result)
				if err != nil {
					return nil, schema.NewDiagnosticsErrorColumnValueExtractor(task.Table, column, err)
				}
				extractor := column_value_extractor.StructSelector("RSA_PUBLIC_KEY")
				return extractor.Extract(ctx, clientMeta, taskClient, task, row, column, r)
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("account").ColumnType(schema.ColumnTypeString).Description("The Snowflake account ID.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {
				return taskClient.(*snowflake_client.Client).Config.Account, nil
			})).Build(),
	}
}

func (x *TableSnowflakeUserGenerator) GetSubTables() []*schema.Table {
	return []*schema.Table{
		table_schema_generator.GenTableSchema(&TableSnowflakeUserGrantGenerator{}),
	}
}
