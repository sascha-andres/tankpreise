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

import "errors"

// GetE10 returns the price for E10
func (sp *StationPrice) GetE10() (float64, error) {
	if !sp.HasE10() {
		return 0.0, errors.New("no such type")
	}
	return sp.E10.(float64), nil
}
