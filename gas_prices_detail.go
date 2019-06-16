package tankpreise

import (
	"encoding/json"
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
