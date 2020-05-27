package models

// MockupConfig ...
type MockupConfig struct {
	ID               int     `json:"id" storm:"increment"`
	URL              string  `json:"url" storm:"unique"`
	Method           string  `json:"method"`
	FailedRatio      float64 `json:"failedRatio"`
	FailedStatusCode int     `json:"failedStatusCode"`
	DataModel        string  `json:"dataModel"`
}
