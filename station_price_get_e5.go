package tankpreise

import "errors"

// GetE5 returns the price for E5
func (sp *StationPrice) GetE5() (float64, error) {
	if !sp.HasE5() {
		return 0.0, errors.New("no such type")
	}
	return sp.E5.(float64), nil
}
