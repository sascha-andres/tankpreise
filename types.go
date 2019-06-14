package tankpreise

/*import (
	"encoding/json"
	"fmt"

	"gopkg.in/resty.v1"
)

type (
	base struct {
		OK bool
	}

	radiusSearch struct {
		base
		License string
	}
)

func main() {
	resp, err := resty.R().Get("https://creativecommons.tankerkoenig.de/json/list.php?lat=52.521&lng=13.438&rad=1.5&sort=dist&type=all&apikey=")
	if err != nil {
		panic(err)
	}
	var s radiusSearch
	err = json.Unmarshal(resp.Body(), &s)
	if err != nil {
		panic(err)
	}
	fmt.Println(s.License)
}*/

// DefaultBaseUrl is the URL where the api is hosted
const DefaultBaseUrl = "https://creativecommons.tankerkoenig.de/json/"

// DemoAPIKey is a key when used will return only dample data and no real data
const DemoAPIKey = "00000000-0000-0000-0000-000000000002"

type (
	// GasPrices is the main entrypoint of the library
	GasPrices struct {
		apiKey      string // apiKey is the key to be used to query the api
		apiEndpoint string // apiEndpoint is the api base url to use
	}

	// BaseResponse contains all properties returned by every API call
	BaseResponse struct {
		OK      bool   // OK denotes success of api call
		Message string // potential error message
	}

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

	// PricesRequest request to get prices for stations listed by id
	PricesRequest struct {
		IDs []string
	}

	// pricesRequest is the translated request object
	pricesRequest struct {
		IDs string `url:"ids"`
	}

	// StationPrice contains prices for gas types
	StationPrice struct {
		Status string      `json:"status"`
		E5     interface{} `json:"e5"`
		E10    interface{} `json:"e10"`
		Diesel interface{} `json:"diesel"`
	}

	// PricesResponse returns prices for gas stations
	PricesResponse struct {
		BaseResponse
		License string                  `json:"license"`
		Data    string                  `json:"data"`
		Prices  map[string]StationPrice `json:"prices"`
	}

	// Station describes a gas station
	Station struct {
		ID          string  `json:"id"`
		Name        string  `json:"name"`
		Brand       string  `json:"brand"`
		Street      string  `json:"street"`
		Place       string  `json:"place"`
		Lat         float64 `json:"lat"`
		Lng         float64 `json:"lng"`
		Dist        float64 `json:"dist"`
		Diesel      float64 `json:"diesel"`
		E5          float64 `json:"e5"`
		E10         float64 `json:"e10"`
		IsOpen      bool    `json:"isOpen"`
		HouseNumber string  `json:"houseNumber"`
		PostCode    int     `json:"postCode"`
	}
)

// NewGasPrices returns a new instance
func NewGasPrices() (*GasPrices, error) {
	return &GasPrices{
		apiKey:      DemoAPIKey,
		apiEndpoint: DefaultBaseUrl,
	}, nil
}
