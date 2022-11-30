package snowflake_client

type Client struct {
	Config *Config
}

func NewClients(configs Configs) ([]*Client, error) {
	var clients []*Client
	for i := range configs.Providers {
		clients = append(clients, &Client{Config: &configs.Providers[i]})
	}
	return clients, nil
}
