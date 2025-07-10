package graphlib


import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
)

type geoNode struct {
	ID 			string
	name 		string
	metadata 	map[string]interface{}
}

type NodePair struct {
	NodeID_from 	string
	NodeID_to 		string
	metaData 		map[string]interface{}
}

type geoEdge struct {
	ConnectingNodes 	NodePair
	descritiondata 		map[string]interface{}
}

type geoGraph struct {
	GraphType 		string
	NodesList 		map[string]geoNode
	ComputedEdges	map[string]geoEdge
	EdgeDelimiter 	string
}

func BuildGeoGraph(ctx context.Context, db *sql.DB, graphID string) (geoGraph, error) {
	query := `
		WITH RECURSIVE region_tree AS (
			SELECT id, name, parent_id, metadata
			FROM regions
			WHERE id = $1
			UNION ALL
			SELECT r.id, r.name, r.parent_id, r.metadata
			FROM regions r
			INNER JOIN region_tree rt ON r.parent_id = rt.id
		)
		SELECT id, name, metadata FROM region_tree;
	`

	rows, err := db.QueryContext(ctx, query, graphID)
	if err != nil {
		return geoGraph{}, fmt.Errorf("query failed: %w", err)
	}
	defer rows.Close()

	graph := geoGraph{}

	for rows.Next() {
		var id, name string
		var metaRaw []byte

		err := rows.Scan(&id, &name, &metaRaw)
		if err != nil {
			return graph, fmt.Errorf("row scan failed: %w", err)
		}

		var meta map[string]interface{}
		if err := json.Unmarshal(metaRaw, &meta); err != nil {
			return graph, fmt.Errorf("json decode failed for id %s: %w", id, err)
		}

		node := geoNode{id, name, meta}
		graph.NodesList[id] = node
	}

	if err := rows.Err(); err != nil {
		return graph, fmt.Errorf("rows error: %w", err)
	}

	return graph, nil
}








