package db

import (
	"database/sql"
	"github.com/MandaKausthubh/GraphLibrary/internal/graph"
)

type DB struct {
	Conn *sql.DB
}

type NodeRepository interface {
	CreateNode(node *graph.Node) (error)
	GetNodeByID(id string) (*graph.Node, error)
	GetChildNodes(parentID string) ([]*graph.Node, error)
	GetMetaData(nodeID string) (map[string]interface{}, error)
}

type EdgeRepository interface {
	CreateEdge(edge *graph.Edge) (error)
	GetEdge(fromID, toID string) (*graph.Edge, error)
	GetEdgesByNodeID(nodeID string) ([]*graph.Edge, error)
}


