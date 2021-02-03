package nodes

type WebhookNode struct {
	name        string
	properties  map[string]interface{}
	credentials map[string]interface{}
}

func (n *WebhookNode) Execute(input []map[string]interface{}) ([]map[string]interface{}, error) {
	return nil, nil
}
