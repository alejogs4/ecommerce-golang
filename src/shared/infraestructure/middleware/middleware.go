package middleware

import "net/http"

// Middleware definition to chain different responsability nodes in http request
type Middleware func(http.HandlerFunc) http.HandlerFunc

// Chain reduce several http handlers into a single one in order to compose several http level behaviors
func Chain(fn http.HandlerFunc, middlewares ...Middleware) http.HandlerFunc {
	for _, m := range middlewares {
		fn = m(fn)
	}

	return fn
}
