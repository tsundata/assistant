package nodes

import (
	"errors"
	"github.com/tsundata/assistant/api/pb"
)

type Node interface {
	Execute(properties map[string]interface{}, credentials map[string]interface{}, input []map[string]interface{}) ([]map[string]interface{}, error)
}

func Execute(name string, regular string, parameters map[string]interface{}, secret string, input []map[string]interface{}, midClient pb.MiddleClient) ([]map[string]interface{}, error) {
	var node Node

	switch regular {
	case "http":
		node = HttpNode{name: name}
	case "cron":
		node = CronNode{name: name}
	default:
		return nil, errors.New("node name error: " + regular)
	}

	/*
		reply, err := midClient.GetCredential(context.Background(), &pb.Text{
			Text: secret,
		})
		if err != nil {
			return "", err
		}
		reply.GetText()
	*/

	return node.Execute(parameters, nil, input)
}
