package models

// MockupConfig ...
type MockupConfig struct {
	ID               string  `json:"id"`
	URL              string  `json:"url"`
	Method           string  `json:"method"`
	FailedRatio      float64 `json:"failedRatio"`
	FailedStatusCode int32   `json:"failedStatusCode"`
	DataModel        string  `json:"dataModel"`
}
