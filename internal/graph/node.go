package graph

import "github.com/google/uuid"

type Node struct {
	ID         string  `json:"id,omitempty"`
	Name       string  `json:"name"`
	Type       string  `json:"type"`
	ParentID   *string `json:"parent_id,omitempty"`
	Latitude   float64 `json:"latitude"`
	Longitude  float64 `json:"longitude"`
	HasPort    bool    `json:"has_port"`
	HasAirport bool    `json:"has_airport"`
	Capacity   int     `json:"capacity"`
}

func (n *Node) IsRoot() bool {
	return n.ParentID == nil
}

func (n *Node) Location() (float64, float64) {
	return n.Latitude, n.Longitude
}

func IsValidUUID(id string) bool {
	_, err := uuid.Parse(id)
	return err == nil
}




