package nodes

import (
	"context"
	"errors"
	"github.com/tsundata/assistant/api/pb"
)

type Node interface {
	Execute(input []map[string]interface{}) ([]map[string]interface{}, error)
}

type Actuator struct {
	midClient pb.MiddleClient

	name       string
	regular    string
	parameters map[string]interface{}
	secret     string
}

func Construct(midClient pb.MiddleClient, name string, regular string, parameters map[string]interface{}, secret string) *Actuator {
	return &Actuator{midClient, name, regular, parameters, secret}
}

func (a *Actuator) Execute(input []map[string]interface{}) ([]map[string]interface{}, error) {
	// credentials
	credentials := make(map[string]interface{})
	if a.midClient != nil {
		reply, err := a.midClient.GetCredentials(context.Background(), &pb.TextRequest{
			Text: a.secret,
		})
		if err != nil {
			return nil, err
		}

		for _, kv := range reply.GetItems() {
			credentials[kv.Key] = kv.Value
		}
	}

	var node Node
	switch a.regular {
	case "http":
		node = &HttpNode{a.name, a.parameters, credentials}
	case "cron":
		node = &CronNode{a.name, a.parameters, credentials}
	default:
		return nil, errors.New("node name error: " + a.regular)
	}

	return node.Execute(input)
}
