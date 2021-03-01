package nodes

type ExecuteNode struct {
	name        string
	properties  map[string]interface{}
	credentials map[string]interface{}
}

func (n *ExecuteNode) Execute(input interface{}) (interface{}, error) {
	return nil, nil
}
