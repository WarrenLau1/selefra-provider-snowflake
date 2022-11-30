package snowflake_client

type Configs struct {
	Providers []Config `yaml:"providers"  mapstructure:"providers"`
}

// Config defines Provider Configuration
type Config struct {
	Account              string `yaml:"account,omitempty" mapstructure:"account"`
	User                 string `yaml:"user,omitempty" mapstructure:"user"`
	Region               string `yaml:"region,omitempty" mapstructure:"region"`
	Role                 string `yaml:"role,omitempty" mapstructure:"role"`
	Password             string `yaml:"password,omitempty" mapstructure:"password"`
	BrowserAuth          bool   `yaml:"browser_auth,omitempty" mapstructure:"browser_auth"`
	PrivateKeyPath       string `yaml:"private_key_path,omitempty" mapstructure:"private_key_path"`
	PrivateKey           string `yaml:"private_key,omitempty" mapstructure:"private_key"`
	PrivateKeyPassphrase string `yaml:"private_key_passphrase,omitempty" mapstructure:"private_key_passphrase"`
	OAuthAccessToken     string `yaml:"oauth_access_token,omitempty" mapstructure:"oauth_access_token"`
	OAuthClientID        string `yaml:"oauth_client_id,omitempty" mapstructure:"oauth_client_id"`
	OAuthClientSecret    string `yaml:"oauth_client_secret,omitempty" mapstructure:"oauth_client_secret"`
	OAuthEndpoint        string `yaml:"oauth_endpoint,omitempty" mapstructure:"oauth_endpoint"`
	OAuthRedirectURL     string `yaml:"oauth_redirect_url,omitempty" mapstructure:"oauth_redirect_url"`
	OAuthRefreshToken    string `yaml:"oauth_refresh_token,omitempty" mapstructure:"oauth_refresh_token"`
	Warehouse            string `yaml:"warehouse,omitempty" mapstructure:"warehouse"`
}
