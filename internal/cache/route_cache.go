package cache


type EdgeCacheEntry struct {
	DistanceKM  	float64   	`json:"distance_km"`
	TravelTimeSec 	int 		`json:"travel_time_sec"`
}


type Cache interface {
	GetEdge(fromID, toID string) (*EdgeCacheEntry, error)
	SetEdge(fromID, toID string, entry *EdgeCacheEntry) error
}
