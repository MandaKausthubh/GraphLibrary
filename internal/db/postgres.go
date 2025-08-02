package db

import (
	"context"
	"log"
	"os"
	"time"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/MandaKausthubh/GraphLibrary/internal/graph"
)


var Pool *pgxpool.Pool

func InitPostgres() {
	url := os.Getenv("POSTGRES_URL")
	config, err := pgxpool.ParseConfig(url)
	if err != nil {
		panic("Failed to parse Postgres URL: " + err.Error())
	}

	config.MaxConns = 10
	config.MinConns = 2
	config.MaxConnIdleTime = time.Hour

	Pool, err = pgxpool.NewWithConfig(context.Background(), config)
	if err != nil {
		panic("Failed to connect to Postgres: " + err.Error())
	}
	if err := Pool.Ping(context.Background()); err != nil {
		panic("Failed to ping Postgres: " + err.Error())
	}
	log.Println("🆗: Connected to Postgres successfully")
}

func GetNodeByID(nodeID string) (*Node, error) {
	var node Node
	err := Pool.QueryRow(context.Background(), "SELECT id, data FROM nodes WHERE id = $1", nodeID).Scan(&node.ID, &node.Data)
	if err != nil {
		return nil, err
	}
	return &node, nil
}

func GetChildNodes(parentID string) ([]*Node, error) {
	rows, err := Pool.Query(context.Background(), "SELECT id, data FROM nodes WHERE parent_id = $1", parentID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var nodes []*Node
	for rows.Next() {
		var node Node
		if err := rows.Scan(&node.ID, &node.Data); err != nil {
			return nil, err
		}
		nodes = append(nodes, &node)
	}
	return nodes, nil
}


func CreateNode(node *Node) error {
	_, err := Pool.Exec(context.Background(), "INSERT INTO nodes (id, data) VALUES ($1, $2)", node.ID, node.Data)
	return err
}

func StoreMetadata(nodeID string, metadata map[string]interface{}) error {
	_, err := Pool.Exec(context.Background(), "INSERT INTO metadata (node_id, data) VALUES ($1, $2)", nodeID, metadata)
	return err
}
