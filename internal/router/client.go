package router

import (
	"fmt"
	"bytes"
	"encoding/json"
	"net/http"
)

func CallGraphHopper(start, end GHPoint, apiKey string) (*GHResponse, error) {
	url := fmt.Sprintf("https://graphhopper.com/api/1/route?key=%s", apiKey)

	reqBody := GHRequest{
		Points:        []GHPoint{start, end},
		Profile:       "car",
		Locale:        "en",
		Instructions:  false,
		PointsEncoded: false,
	}

	body, _ := json.Marshal(reqBody)
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var ghResp GHResponse
	if err := json.NewDecoder(resp.Body).Decode(&ghResp); err != nil {
		return nil, err
	}

	return &ghResp, nil
}









