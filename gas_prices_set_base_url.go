package tankpreise

// SetBaseUrl can be used to override default api base
func (gp *GasPrices) SetBaseUrl(url string) {
	gp.apiEndpoint = url
}
