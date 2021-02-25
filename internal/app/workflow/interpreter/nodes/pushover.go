package nodes

import (
	"github.com/go-resty/resty/v2"
	"time"
)

type PushoverNode struct {
	name        string
	properties  map[string]interface{}
	credentials map[string]interface{}
}

func (n *PushoverNode) Execute(input interface{}) (interface{}, error) {
	token := extractCredentials(n.credentials, "token")
	user := extractCredentials(n.credentials, "user")

	title := extractProperties(input, n.properties, "title")
	message := extractProperties(input, n.properties, "message")
	device := extractProperties(input, n.properties, "device")
	url := extractProperties(input, n.properties, "url")
	urlTitle := extractProperties(input, n.properties, "url_title")
	priority := extractProperties(input, n.properties, "priority")
	timestamp := extractProperties(input, n.properties, "timestamp")

	client := resty.New()
	client.SetTimeout(time.Minute)

	resp, err := client.R().
		SetResult(&map[string]interface{}{}).
		SetBody(map[string]interface{}{
			"token":   token,
			"user":    user,
			"message": message,
			// optional
			"title":     title,
			"device":    device,
			"url":       url,
			"url_title": urlTitle,
			"priority":  priority, // -2,-1,1,2
			"timestamp": timestamp,
		}).
		Post("https://api.pushover.net/1/messages.json")

	if err != nil {
		return nil, err
	}

	if res, ok := resp.Result().(*map[string]interface{}); ok {
		return *res, nil
	}

	return nil, nil
}
