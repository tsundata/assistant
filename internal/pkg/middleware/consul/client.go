package consul

import (
	"github.com/google/wire"
	"github.com/hashicorp/consul/api"
	"os"
)

func New() (*api.Client, error) {
	consulAddress := os.Getenv("CONSUL_ADDRESS")
	client, err := api.NewClient(&api.Config{
		Address: consulAddress,
		Scheme:  "http", // todo
	})
	if err != nil {
		return nil, err
	}
	return client, nil
}

var ProviderSet = wire.NewSet(New)
