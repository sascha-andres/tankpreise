package tankpreise

// HasE% returns true if the station has type E5
func (sp *StationPrice) HasE10() bool {
	if sp.E10 == nil {
		return false
	}
	switch sp.E10.(type) {
	case float64:
		return true
	}
	return false
}
