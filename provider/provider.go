package provider

import (
	"context"
	"github.com/selefra/selefra-provider-snowflake/snowflake_client"
	"os"

	"github.com/selefra/selefra-provider-sdk/provider"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/spf13/viper"
)

const Version = "v0.0.1"

func GetProvider() *provider.Provider {
	return &provider.Provider{
		Name:      "snowflake",
		Version:   Version,
		TableList: GenTables(),
		ClientMeta: schema.ClientMeta{
			InitClient: func(ctx context.Context, clientMeta *schema.ClientMeta, config *viper.Viper) ([]any, *schema.Diagnostics) {
				var snowflakeConfig snowflake_client.Configs

				err := config.Unmarshal(&snowflakeConfig.Providers)
				if err != nil {
					return nil, schema.NewDiagnostics().AddErrorMsg("analysis config err: %s", err.Error())
				}

				if len(snowflakeConfig.Providers) == 0 {
					snowflakeConfig.Providers = append(snowflakeConfig.Providers, snowflake_client.Config{})
				}

				if snowflakeConfig.Providers[0].Account == "" {
					snowflakeConfig.Providers[0].Account = os.Getenv("SNOWFLAKE_ACCOUNT")
				}

				if snowflakeConfig.Providers[0].Account == "" {
					return nil, schema.NewDiagnostics().AddErrorMsg("missing Account in configuration")
				}

				if snowflakeConfig.Providers[0].Role == "" {
					snowflakeConfig.Providers[0].Role = os.Getenv("SNOWFLAKE_ROLE")
				}

				if snowflakeConfig.Providers[0].Role == "" {
					return nil, schema.NewDiagnostics().AddErrorMsg("missing Role in configuration")
				}

				if snowflakeConfig.Providers[0].User == "" {
					snowflakeConfig.Providers[0].User = os.Getenv("SNOWFLAKE_USER")
				}

				if snowflakeConfig.Providers[0].User == "" {
					return nil, schema.NewDiagnostics().AddErrorMsg("missing User in configuration")
				}

				if snowflakeConfig.Providers[0].Password == "" {
					snowflakeConfig.Providers[0].Password = os.Getenv("SNOWFLAKE_PASSWORD")
				}

				if snowflakeConfig.Providers[0].Password == "" {
					return nil, schema.NewDiagnostics().AddErrorMsg("missing Password in configuration")
				}

				if snowflakeConfig.Providers[0].Warehouse == "" {
					snowflakeConfig.Providers[0].Warehouse = os.Getenv("SNOWFLAKE_WAREHOUSE")
				}

				if snowflakeConfig.Providers[0].Warehouse == "" {
					return nil, schema.NewDiagnostics().AddErrorMsg("missing Warehouse in configuration")
				}

				clients, err := snowflake_client.NewClients(snowflakeConfig)

				if err != nil {
					clientMeta.ErrorF("new clients err: %s", err.Error())
					return nil, schema.NewDiagnostics().AddError(err)
				}

				if len(clients) == 0 {
					return nil, schema.NewDiagnostics().AddErrorMsg("account information not found")
				}

				res := make([]interface{}, 0, len(clients))
				for i := range clients {
					res = append(res, clients[i])
				}
				return res, nil
			},
		},
		ConfigMeta: provider.ConfigMeta{
			GetDefaultConfigTemplate: func(ctx context.Context) string {
				return `# account: "<YOUR_ACCOUNT>" 
# role: "<YOUR_ROLE>"
# user: "<YOUR_USERNAME>"
# password: "<YOUR_USER_PASSWORD>"
# warehouse: "<YOUR_WAREHOUSE>"
`
			},
			Validation: func(ctx context.Context, config *viper.Viper) *schema.Diagnostics {
				var snowflakeConfig snowflake_client.Configs
				err := config.Unmarshal(&snowflakeConfig.Providers)

				if err != nil {
					return schema.NewDiagnostics().AddErrorMsg("analysis config err: %s", err.Error())
				}

				return nil
			},
		},
		TransformerMeta: schema.TransformerMeta{
			DefaultColumnValueConvertorBlackList: []string{
				"",
				"N/A",
				"not_supported",
			},
			DataSourcePullResultAutoExpand: true,
		},
		ErrorsHandlerMeta: schema.ErrorsHandlerMeta{

			IgnoredErrors: []schema.IgnoredError{schema.IgnoredErrorOnSaveResult},
		},
	}
}
