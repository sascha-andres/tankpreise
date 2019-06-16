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

	// Station describes a gas station
	Station struct {
		ID           string `json:"id"`
		Name         string `json:"name"`
		Brand        string `json:"brand"`
		Street       string `json:"street"`
		HouseNumber  string `json:"houseNumber"`
		PostCode     int    `json:"postCode"`
		Place        string `json:"place"`
		OpeningTimes []struct {
			Text  string `json:"text"`
			Start string `json:"start"`
			End   string `json:"end"`
		} `json:"openingTimes"`
		Overrides []string `json:"overrides"`
		WholeDay  bool     `json:"wholeDay"`
		IsOpen    bool     `json:"isOpen"`
		E5        float64  `json:"e5"`
		E10       float64  `json:"e10"`
		Diesel    float64  `json:"diesel"`
		Lat       float64  `json:"lat"`
		Lng       float64  `json:"lng"`
		Dist      float64  `json:"dist"`
		State     string   `json:"state"`
	}
)

// NewGasPrices returns a new instance
func NewGasPrices() (*GasPrices, error) {
	return &GasPrices{
		apiKey:      DemoAPIKey,
		apiEndpoint: DefaultBaseUrl,
	}, nil
}
