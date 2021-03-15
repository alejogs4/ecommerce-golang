package httputils

// HTTPError struct to modeled a standard http error
type HTTPError struct {
	Message    string
	StatusCode int
}
