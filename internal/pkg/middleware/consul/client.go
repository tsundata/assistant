package consul

import (
	"github.com/google/wire"
	"github.com/hashicorp/consul/api"
	"os"
)

func New() (*api.Client, error) {
	consulAddress := os.Getenv("CONSUL_ADDRESS")
	consulScheme := os.Getenv("CONSUL_SCHEME")
	consulUsername := os.Getenv("CONSUL_USERNAME")
	consulPassword := os.Getenv("CONSUL_PASSWORD")
	consulToken := os.Getenv("CONSUL_TOKEN")
	client, err := api.NewClient(&api.Config{
		Address: consulAddress,
		Scheme:  consulScheme,
		HttpAuth: &api.HttpBasicAuth{
			Username: consulUsername,
			Password: consulPassword,
		},
		Token: consulToken,
	})
	if err != nil {
		return nil, err
	}
	return client, nil
}

var ProviderSet = wire.NewSet(New)
