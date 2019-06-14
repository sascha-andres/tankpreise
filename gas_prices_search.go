package tankpreise

import (
	"encoding/json"
)

func (gp *GasPrices) Search(query SearchRequest) (*SearchResponse, error) {
	body, err := gp.do("GET", "list.php", query)
	if err != nil {
		return nil, err
	}
	var s SearchResponse
	err = json.Unmarshal(body, &s)
	return &s, nil
}
