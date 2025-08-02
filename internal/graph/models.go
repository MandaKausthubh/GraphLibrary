package graph

import (
	"time"
	"github.com/google/uuid"
	"fmt"
	"context"
	"github.com/jmoiron/sqlx"
)

type Node struct {
    ID           uuid.UUID
    Name         string
    Type         string // country, state, city, etc.
    ParentID     *uuid.UUID
    Latitude     float64
    Longitude    float64
    HasSeaport   bool
    HasAirport   bool
    Capacity     int
    CreatedAt    time.Time
    Metadata     map[string]interface{}
}

type Edge struct {
    ID           uuid.UUID
    FromNodeID   uuid.UUID
    ToNodeID     uuid.UUID
    DistanceKM   float64
    TravelTimeS  int
    Metadata     map[string]interface{}
    CreatedAt    time.Time
}

type Graph struct {
    Nodes map[uuid.UUID]*Node
    Edges []*Edge
}


