package nodes

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/tidwall/gjson"
	"github.com/tsundata/assistant/api/pb"
	"github.com/tsundata/assistant/internal/pkg/utils"
	"regexp"
	"strings"
)

type Node interface {
	Execute(input interface{}) (interface{}, error)
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

func (a *Actuator) Execute(input interface{}) (interface{}, error) {
	// credentials
	credentials := make(map[string]interface{})
	if a.midClient != nil && a.secret != "" {
		reply, err := a.midClient.GetCredential(context.Background(), &pb.CredentialRequest{
			Name: a.secret,
		})
		if err != nil {
			return nil, err
		}

		for _, kv := range reply.GetContent() {
			credentials[kv.Key] = kv.Value
		}
	}

	var node Node
	switch a.regular {
	case "http":
		node = &HttpNode{a.name, a.parameters, credentials}
	case "cron":
		node = &CronNode{a.name, a.parameters, credentials}
	case "webhook":
		node = &WebhookNode{a.name, a.parameters, credentials}
	case "execute":
		node = &ExecuteNode{a.name, a.parameters, credentials}
	case "pushover":
		node = &PushoverNode{a.name, a.parameters, credentials}
	default:
		return nil, errors.New("node name error: " + a.regular)
	}

	return node.Execute(input)
}

func extractCredentials(credentials map[string]interface{}, key string) interface{} {
	return credentials[key]
}

func extractProperties(input interface{}, properties map[string]interface{}, key string) interface{} {
	data, ok := properties[key]
	if !ok {
		return nil
	}
	value, ok := data.(string)
	if !ok {
		return data
	}

	inputJson, err := json.Marshal(input)
	if err != nil {
		return nil
	}

	// expression
	re := regexp.MustCompile(`\{\{[a-zA-Z0-9.\*#]+\}\}`)
	expression := re.FindAllString(value, -1)
	for _, e := range expression {
		path := strings.ReplaceAll(e, "{{", "")
		path = strings.ReplaceAll(path, "}}", "")
		result := gjson.Get(utils.ByteToString(inputJson), path)
		value = strings.ReplaceAll(value, e, result.String())
	}

	return value
}
