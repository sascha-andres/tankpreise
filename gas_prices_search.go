package tankpreise

import (
	"encoding/json"
)

type (
	// SearchRequest lets you provide all parameters for a radius search
	SearchRequest struct {
		Latitude  float64 `url:"lat"`  // geographic width of location
		Longitude float64 `url:"lng"`  // geographic height of location
		Radius    float64 `url:"rad"`  // Radius to search within
		GasType   string  `url:"type"` // GasType denotes wich type of gas to lopok for (one of e5, e10, diesel or all)
		Sort      string  `url:"sort"` // Sort allows to specify whether to search by dist or price, no effect when GasType == all
	}

	// SearchResponse contains all data returned by a radius search
	SearchResponse struct {
		BaseResponse
		License  string    `json:"license"`
		Data     string    `json:"data"`
		Status   string    `json:"status"`
		Stations []Station `json:"stations"`
	}
)

// Search returns a list of stations
func (gp *GasPrices) Search(query SearchRequest) (*SearchResponse, error) {
	body, err := gp.do("GET", "list.php", query)
	if err != nil {
		return nil, err
	}
	var s SearchResponse
	err = json.Unmarshal(body, &s)
	return &s, nil
}
