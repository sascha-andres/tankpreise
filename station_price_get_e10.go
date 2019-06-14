package tankpreise

import "errors"

// GetE10 returns the price for E10
func (sp *StationPrice) GetE10() (float64, error) {
	if !sp.HasE10() {
		return 0.0, errors.New("no such type")
	}
	return sp.E10.(float64), nil
}
