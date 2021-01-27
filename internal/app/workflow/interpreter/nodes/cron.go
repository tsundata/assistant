package nodes

type CronNode struct {
	name string
}

func (n CronNode) Execute(properties map[string]interface{}, credentials map[string]interface{}, input string) (string, error) {
	return "", nil
}
