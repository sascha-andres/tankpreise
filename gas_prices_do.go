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
		if err == nil {
			v.Add("apikey", gp.apiKey)
			url := strings.TrimSpace(fmt.Sprintf("%s%s?%s\n", gp.apiEndpoint, endpoint, v.Encode()))
			resp, err = resty.R().Get(url)
		}
	default:
		err = errors.New("called with unsupported method")
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
