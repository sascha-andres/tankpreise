/*
Copyright © 2019 Sascha Andres <sascha.andres@outlook.com>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

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
