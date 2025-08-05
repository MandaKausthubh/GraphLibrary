package db

import (
	"database/sql"
	"github.com/MandaKausthubh/GraphLibrary/internal/graph"
)

type EdgeRepositoryImpl struct {
	DB *sql.DB
}

func (r *EdgeRepositoryImpl) CreateEdge(edge *graph.Edge) error {
	query := `
		INSERT INTO edges (edge_id, from_node_id, to_node_id, distance_km, travel_time_sec, metadata)
		VALUES ($1, $2, $3, $4, $5, $6)
	`
	_, err := r.DB.Exec(query, edge.ID, edge.FromNodeID, edge.ToNodeID, edge.DistanceKM, edge.TravelTimeSec, edge.Metadata)
	return err
}

func (r *EdgeRepositoryImpl) GetEdge(fromID, toID string) (*graph.Edge, error) {
	query := `
		SELECT edge_id, from_node_id, to_node_id, distance_km, travel_time_sec, metadata
		FROM edges
		WHERE from_node_id = $1 AND to_node_id = $2
	`
	row := r.DB.QueryRow(query, fromID, toID)

	var edge graph.Edge
	err := row.Scan(&edge.ID, &edge.FromNodeID, &edge.ToNodeID, &edge.DistanceKM, &edge.TravelTimeSec, &edge.Metadata)
	if err != nil {
		return nil, err
	}
	return &edge, nil
}

func (r *EdgeRepositoryImpl) GetEdgesByNodeID(nodeID string) ([]*graph.Edge, error) {
	query := `
		SELECT edge_id, from_node_id, to_node_id, distance_km, travel_time_sec, metadata
		FROM edges
		WHERE from_node_id = $1 OR to_node_id = $1
	`
	rows, err := r.DB.Query(query, nodeID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var edges []*graph.Edge
	for rows.Next() {
		var edge graph.Edge
		if err := rows.Scan(&edge.ID, &edge.FromNodeID, &edge.ToNodeID, &edge.DistanceKM, &edge.TravelTimeSec, &edge.Metadata); err != nil {
			return nil, err
		}
		edges = append(edges, &edge)
	}
	return edges, nil
}
