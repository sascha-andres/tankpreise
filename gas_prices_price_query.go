package tankpreise

import "encoding/json"

type (
	// PricesRequest request to get prices for stations listed by id
	PricesRequest struct {
		IDs []string
	}

	// pricesRequest is the translated request object
	pricesRequest struct {
		IDs string `url:"ids"`
	}

	// PricesResponse returns prices for gas stations
	PricesResponse struct {
		BaseResponse
		License string                  `json:"license"`
		Data    string                  `json:"data"`
		Prices  map[string]StationPrice `json:"prices"`
	}

	// StationPrice contains prices for gas types
	StationPrice struct {
		Status string      `json:"status"`
		E5     interface{} `json:"e5"`
		E10    interface{} `json:"e10"`
		Diesel interface{} `json:"diesel"`
	}
)

// PriceQuery returns prices found
func (gp *GasPrices) PriceQuery(query PricesRequest) (*PricesResponse, error) {
	q := pricesRequest{
		IDs: query.String(),
	}
	body, err := gp.do("GET", "prices.php", q)
	if err != nil {
		return nil, err
	}
	var pr PricesResponse
	err = json.Unmarshal(body, &pr)
	return &pr, err
}
