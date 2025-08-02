package graph

import (
	"github.com/google/uuid"
	"fmt"
	"context"
	"github.com/jmoiron/sqlx"
)

func (g *Graph) AddNode(node *Node) int {
	for _, existingNode := range g.Nodes {
		if existingNode.ID == node.ID {
			return 0
		}
	}
	g.Nodes[node.ID] = node
	return 1
}

func (g *Graph) AddEdge(edge *Edge) int {
	for _, existingEdge := range g.Edges {
		if existingEdge.ID == edge.ID {
			return 0
		}
	}
	g.Edges = append(g.Edges, edge)
	return 1
}

func (g *Graph) ImmediateSubgraph(ctx context.Context, db *sqlx.DB, rootID uuid.UUID) (*Graph, error) {
    subgraph := &Graph{
        Nodes: make(map[uuid.UUID]*Node),
        Edges: make([]*Edge, 0),
    }

    var nodes []*Node
    query := `
        SELECT * FROM nodes WHERE parent_node_id = $1;
    `
    if err := db.SelectContext(ctx, &nodes, query, rootID); err != nil {
        return nil, fmt.Errorf("failed to fetch children: %w", err)
    }

    var root Node
    if err := db.GetContext(ctx, &root, `SELECT * FROM nodes WHERE node_id = $1`, rootID); err != nil {
        return nil, fmt.Errorf("failed to fetch root node: %w", err)
    }

    subgraph.Nodes[root.ID] = &root
    for _, n := range nodes {
        subgraph.Nodes[n.ID] = n
    }

    nodeIDs := make([]interface{}, 0, len(subgraph.Nodes))
    for id := range subgraph.Nodes {
        nodeIDs = append(nodeIDs, id)
    }

    edgeQuery, args, err := sqlx.In(`
        SELECT * FROM edges
        WHERE from_node_id IN (?) AND to_node_id IN (?)
    `, nodeIDs, nodeIDs)
    if err != nil {
        return nil, err
    }
    edgeQuery = db.Rebind(edgeQuery)

    var edges []*Edge
    if err := db.SelectContext(ctx, &edges, edgeQuery, args...); err != nil {
        return nil, err
    }

    subgraph.Edges = edges

    return subgraph, nil
}





