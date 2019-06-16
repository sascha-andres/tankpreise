package tankpreise

import (
	"encoding/json"
)

type (
	// DetailRequest is used to get details about a gas station
	DetailRequest struct {
		ID string `url:"id"`
	}

	// DetailResponse contains data returned for a detail request
	DetailResponse struct {
		BaseResponse
		License string  `json:"license"`
		Data    string  `json:"data"`
		Status  string  `json:"status"`
		Station Station `json:"station"`
	}
)

// Detail returns details about a gas station
func (gp *GasPrices) Detail(query DetailRequest) (*DetailResponse, error) {
	data, err := gp.do("GET", "detail.php", query)
	if err != nil {
		return nil, err
	}
	var result DetailResponse
	err = json.Unmarshal(data, &result)
	return &result, err
}
