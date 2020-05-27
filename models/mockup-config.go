package models

// MockupConfig ...
type MockupConfig struct {
	ID               int     `json:"id" storm:"increment"`
	URL              string  `json:"url"`
	Method           string  `json:"method"`
	FailedRatio      float64 `json:"failedRatio"`
	FailedStatusCode int     `json:"failedStatusCode"`
	DataModel        string  `json:"dataModel"`
}

// ValidateMockupConfig ...
func ValidateMockupConfig(data MockupConfig) error {
	// TODO: validate before create / update

	/**
	1. unique URL + Method
	2. DataModel
		- field -> oneOf [boolean, int, float, string, date (iso)]
		- field -> array -> length = 1
		- field.format -> { min, max, format: oneOf [UPPERCASE, lowercase, snake_case, camelCase, imageUrl] } (must have both min/max or none)
		- field.count -> { min, max } (must have both)
	*/

	return nil
}
