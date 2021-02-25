package nodes

type CronNode struct {
	name        string
	properties  map[string]interface{}
	credentials map[string]interface{}
}

func (n *CronNode) Execute(input interface{}) (interface{}, error) {
	return nil, nil
}
