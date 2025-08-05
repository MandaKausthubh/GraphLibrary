package graph

import (
	"time"
	"github.com/google/uuid"
)

type Edge struct {
	EdgeID        uuid.UUID `json:"edge_id"`
	FromNodeID    uuid.UUID `json:"from_node_id"`
	ToNodeID      uuid.UUID `json:"to_node_id"`
	DistanceKm    float64   `json:"distance_km"`
	TravelTimeSec int       `json:"travel_time_sec"`
	Metadata      string    `json:"metadata"`
	CreatedAt     time.Time `json:"created_at"`
}
