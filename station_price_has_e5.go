package tankpreise

// HasE% returns true if the station has type E5
func (sp *StationPrice) HasE5() bool {
	if sp.E5 == nil {
		return false
	}
	switch sp.E5.(type) {
	case float64:
		return true
	}
	return false
}
