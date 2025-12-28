package omgo

import "fmt"

// APIError represents an error returned by the Open-Meteo API.
type APIError struct {
	StatusCode int
	Reason     string
}

// Error implements the error interface.
func (e *APIError) Error() string {
	return fmt.Sprintf("open-meteo error %d: %s", e.StatusCode, e.Reason)
}

// apiErrorResponse represents the JSON structure of an API error response.
type apiErrorResponse struct {
	Error  bool   `json:"error"`
	Reason string `json:"reason"`
}

