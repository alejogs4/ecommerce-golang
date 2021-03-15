package httputils

import (
	"fmt"
	"net/http"

	"github.com/alejogs4/hn-website/src/shared/infraestructure/middleware"
)

// Method middleware ensure that used method is the specified one, returning to use a 405 if request was made with the wrong one
func Method(method string) middleware.Middleware {
	return func(hf http.HandlerFunc) http.HandlerFunc {
		return func(rw http.ResponseWriter, r *http.Request) {
			rw.Header().Set("Content-type", "application/json")

			if r.Method != method {
				DispatchHTTPError(rw, fmt.Sprintf("Method %s is not allowed for this path", r.Method), http.StatusMethodNotAllowed)
				return
			}

			hf(rw, r)
		}
	}
}
