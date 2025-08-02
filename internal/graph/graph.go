package graph

type Graph struct {
	Nodes []*Node
	Edges []*Edge
}

func (g *Graph) AddNode(node *Node) (int)
func (g *Graph) AddEdge(edge *Edge) (int)




