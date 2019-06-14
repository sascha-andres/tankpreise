package tankpreise

import "errors"

// GetDiesel returns the price for diesel
func (sp *StationPrice) GetDiesel() (float64, error) {
	if !sp.HasDiesel() {
		return 0.0, errors.New("no such type")
	}
	return sp.Diesel.(float64), nil
}
