package tankpreise

// HasDiesel returns true if the station has type diesel
func (sp *StationPrice) HasDiesel() bool {
	if sp.Diesel == nil {
		return false
	}
	switch sp.Diesel.(type) {
	case float64:
		return true
	}
	return false
}
