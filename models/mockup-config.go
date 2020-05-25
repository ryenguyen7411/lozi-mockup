package models

// MockupConfig ...
type MockupConfig struct {
	ID               int     `json:"id"`
	URL              string  `json:"url"`
	Method           string  `json:"method"`
	FailedRatio      float64 `json:"failedRatio"`
	FailedStatusCode int     `json:"failedStatusCode"`
	DataModel        string  `json:"dataModel"`
}
