package graph

import (
	"time"
)

type Edge struct {
	EdgeID        string `json:"edge_id"`
	FromNodeID    string `json:"from_node_id"`
	ToNodeID      string `json:"to_node_id"`
	DistanceKm    float64   `json:"distance_km"`
	TravelTimeSec int       `json:"travel_time_sec"`
	Metadata      string    `json:"metadata"`
	CreatedAt     time.Time `json:"created_at"`
}
