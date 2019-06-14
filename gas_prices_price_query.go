package tankpreise

import "encoding/json"

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
