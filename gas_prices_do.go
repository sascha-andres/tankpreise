package tankpreise

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/google/go-querystring/query"
	"gopkg.in/resty.v1"
	"strings"
)

// do is used to make a http request
func (gp *GasPrices) do(method, endpoint string, parameters interface{}) ([]byte, error) {
	var (
		resp *resty.Response
		err  error
	)
	switch method {
	case "GET":
		v, err := query.Values(parameters)
		if err != nil {
			return nil, err
		}
		v.Add("apikey", gp.apiKey)
		url := strings.TrimSpace(fmt.Sprintf("%s%s?%s\n", gp.apiEndpoint, endpoint, v.Encode()))
		resp, err = resty.R().Get(url)
	default:
		return nil, errors.New("called with unsupported method")
	}
	if err != nil {
		return nil, err
	}
	var s BaseResponse
	err = json.Unmarshal(resp.Body(), &s)
	if err != nil {
		return nil, err
	}
	if !s.OK {
		return nil, errors.New("error: " + s.Message)
	}
	return resp.Body(), nil
}
