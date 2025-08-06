package db

import (
	"database/sql"
	"encoding/json"
	"github.com/MandaKausthubh/GraphLibrary/internal/graph"
)


type NodeRepositoryImpl struct {
	DB *sql.DB
}

func (r *NodeRepositoryImpl) CreateNode(node *graph.Node) error {
	query := `
		INSERT INTO nodes (node_id, name, type, parent_id, latitude, longitude, has_port, has_airport, capacity)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
	`
	_, err := r.DB.Exec(query, node.ID, node.Name, node.Type, node.ParentID, node.Latitude, node.Longitude, node.HasPort, node.HasAirport, node.Capacity)
	return err
}

func (r *NodeRepositoryImpl) GetNodeByID(id string) (*graph.Node, error) {
	query := `SELECT node_id, name, type, parent_id, latitude, longitude, has_port, has_airport, capacity FROM nodes WHERE node_id = $1`
	row := r.DB.QueryRow(query, id)

	var node graph.Node
	err := row.Scan(&node.ID, &node.Name, &node.Type, &node.ParentID, &node.Latitude, &node.Longitude, &node.HasPort, &node.HasAirport, &node.Capacity)
	if err != nil {
		return nil, err
	}
	return &node, nil
}

func (r *NodeRepositoryImpl) GetChildNodes(parentID string) ([]*graph.Node, error) {
	query := `SELECT node_id, name, type, parent_id, latitude, longitude, has_port, has_airport, capacity FROM nodes WHERE parent_id = $1`
	rows, err := r.DB.Query(query, parentID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var nodes []*graph.Node
	for rows.Next() {
		var node graph.Node
		if err := rows.Scan(&node.ID, &node.Name, &node.Type, &node.ParentID, &node.Latitude, &node.Longitude, &node.HasPort, &node.HasAirport, &node.Capacity); err != nil {
			return nil, err
		}
		nodes = append(nodes, &node)
	}
	return nodes, nil
}

func (r *NodeRepositoryImpl) GetMetaData(nodeID string) (map[string]interface{}, error) {
	query := `SELECT metadata FROM nodes WHERE node_id = $1`
	var rawMeta json.RawMessage
	err := r.DB.QueryRow(query, nodeID).Scan(&rawMeta)
	if err != nil {
		return nil, err
	}

	var metadata map[string]interface{}
	if err := json.Unmarshal(rawMeta, &metadata); err != nil {
		return nil, err
	}
	return metadata, nil
}






