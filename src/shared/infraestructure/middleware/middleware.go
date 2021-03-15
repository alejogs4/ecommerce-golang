package middleware

import "net/http"

// Middleware definition to chain different responsability nodes in http request
type Middleware func(http.HandlerFunc) http.HandlerFunc

// Chain reduce several http handlers into a single one in order to compose several http level behaviors
func Chain(fn http.HandlerFunc, middlewares ...Middleware) http.HandlerFunc {
	for i := len(middlewares) - 1; i > -1; i-- {
		fn = middlewares[i](fn)
	}

	return fn
}
