package router

type GHPoint [2]float64

type GHRequest struct {
	Points 			[]GHPoint 	`json:"points"`
	Profile        	string    	`json:"profile"`
	Locale         	string    	`json:"locale"`
	Instructions   	bool      	`json:"instructions"`
	PointsEncoded  	bool      	`json:"points_encoded"`
	CalcPoints 	    bool 	 	`json:"calc_points"`
}

type GHResponse struct {
	Hints map[string]interface{} `json:"hints"`

	Info struct {
		Copyrights         []string `json:"copyrights"`
		Took               int      `json:"took"`
		RoadDataTimestamp  string   `json:"road_data_timestamp"`
	} `json:"info"`

	Paths []struct {
		Distance     float64 `json:"distance"`
		Weight       float64 `json:"weight"`
		Time         int64   `json:"time"`
		Transfers    int     `json:"transfers"`
		PointsEncoded bool   `json:"points_encoded"`
		BBox         []float64 `json:"bbox"`

		Points struct {
			Type        string        `json:"type"`
			Coordinates [][]float64   `json:"coordinates"`
		} `json:"points"`

		Instructions []struct {
			Distance float64 `json:"distance"`
			Heading  float64 `json:"heading"`
			Sign     int     `json:"sign"`
			Interval []int   `json:"interval"`
			Text     string  `json:"text"`
			Time     int64   `json:"time"`
			StreetName string `json:"street_name"`
		} `json:"instructions"`

		Details map[string][][]interface{} `json:"details"`

		Ascend          float64 `json:"ascend"`
		Descend         float64 `json:"descend"`

		SnappedWaypoints struct {
			Type        string        `json:"type"`
			Coordinates [][]float64   `json:"coordinates"`
		} `json:"snapped_waypoints"`
	} `json:"paths"`
}

