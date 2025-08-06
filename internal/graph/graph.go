package graph

type Graph struct {
	Nodes []*Node
	Edges []*Edge
}

// AddNode adds a node if not already present
func (g *Graph) AddNode(node *Node) {
	for _, n := range g.Nodes {
		if n.ID == node.ID {
			return
		}
	}
	g.Nodes = append(g.Nodes, node)
}

// AddEdge adds an edge if not already present
func (g *Graph) AddEdge(edge *Edge) {
	for _, e := range g.Edges {
		if e.FromNodeID == edge.FromNodeID && e.ToNodeID == edge.ToNodeID {
			return
		}
	}
	g.Edges = append(g.Edges, edge)
}

// GetNodeByID returns a node by ID
func (g *Graph) GetNodeByID(id string) *Node {
	for _, n := range g.Nodes {
		if n.ID == id {
			return n
		}
	}
	return nil
}

// GetOutgoingEdges returns all edges from a given node
func (g *Graph) GetOutgoingEdges(fromNodeID string) []*Edge {
	var out []*Edge
	for _, e := range g.Edges {
		if e.FromNodeID == fromNodeID {
			out = append(out, e)
		}
	}
	return out
}

// GetNeighbors returns the nodes reachable from a given node
func (g *Graph) GetNeighbors(nodeID string) []*Node {
	var neighbors []*Node
	for _, edge := range g.Edges {
		if edge.FromNodeID == nodeID {
			n := g.GetNodeByID(edge.ToNodeID)
			if n != nil {
				neighbors = append(neighbors, n)
			}
		}
	}
	return neighbors
}

