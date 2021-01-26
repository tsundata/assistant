package nodes

type HttpNode struct {
	name string
}

func (n HttpNode) Execute(properties map[string]interface{}, credentials map[string]interface{}, input string) string {
	return ""
}
