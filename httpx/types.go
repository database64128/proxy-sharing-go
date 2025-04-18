package httpx

// StandardError is the standard error response.
type StandardError struct {
	Message string `json:"error"`
}
