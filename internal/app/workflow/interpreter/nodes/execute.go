package nodes

type ExecuteNode struct {
	name        string
	properties  map[string]interface{}
	credentials map[string]interface{}
}

func (n *ExecuteNode) Execute(input []map[string]interface{}) ([]map[string]interface{}, error) {
	return nil, nil
}
