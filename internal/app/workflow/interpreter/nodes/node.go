package nodes

type Node interface {
	Execute(properties map[string]interface{}, credentials map[string]interface{}, input string) string
}

func Execute(name string, parameters map[string]interface{}, secret string, input string) string {
	var node Node

	switch name {
	case "http":
		node = HttpNode{}
	}
	if node == nil {
		return ""
	}

	// todo call grpc

	return node.Execute(parameters, nil, input)
}
