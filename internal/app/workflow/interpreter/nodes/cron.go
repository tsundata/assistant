package nodes

type CronNode struct {
	name string
}

func (n CronNode) Execute(properties map[string]interface{}, credentials map[string]interface{}, input []map[string]interface{}) ([]map[string]interface{}, error) {
	return nil, nil
}
