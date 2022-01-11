package automate

type Edge struct {
	FromNode *Node
	ToNode *Node
}

func AddEdge(from *Node, to *Node) *Edge {
	edge := &Edge{
		FromNode: from,
		ToNode:   to,
	}
	from.Children = append(from.Children, edge)
	to.Dependency = append(to.Dependency, edge)
	return edge
}