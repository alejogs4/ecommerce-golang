package httputils

import (
	"encoding/json"
	"net/http"
)

// DispatchHTTPError send to user a http error with a body containing an error message
func DispatchHTTPError(w http.ResponseWriter, message string, statusCode int) {
	content, _ := json.Marshal(map[string]string{"message": message})

	w.WriteHeader(statusCode)
	w.Write(content)
}

// DispatchNewResponse sends any interface as a http response body for a hopefully succesfull operation
func DispatchNewResponse(w http.ResponseWriter, data interface{}, statusCode int) {
	content, _ := json.Marshal(data)

	w.WriteHeader(statusCode)
	w.Write(content)
}

// DispatchDefaultAPIResponse sends a default http response, the same way as DispatchNewResponse does
func DispatchDefaultAPIResponse(w http.ResponseWriter, data interface{}, message string, statusCode int) {
	DispatchNewResponse(w, wrapAPIResponse(data, message), statusCode)
}

func wrapAPIResponse(data interface{}, message string) map[string]interface{} {
	return map[string]interface{}{
		"data":    data,
		"message": message,
	}
}
