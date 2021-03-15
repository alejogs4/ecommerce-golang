package authorization

import (
	"context"
	"net/http"
	"strings"

	httputils "github.com/alejogs4/hn-website/src/shared/infraestructure/httpUtils"
	"github.com/alejogs4/hn-website/src/shared/infraestructure/middleware"
	"github.com/alejogs4/hn-website/src/shared/infraestructure/token"
	userdto "github.com/alejogs4/hn-website/src/user/application/userDTO"
	usererrormapper "github.com/alejogs4/hn-website/src/user/infraestructure/userHttpPort/userErrorMapper"
)

// ValidateToken from http headers
func ValidateToken() middleware.Middleware {
	return func(hf http.HandlerFunc) http.HandlerFunc {
		return func(rw http.ResponseWriter, r *http.Request) {
			rw.Header().Set("Content-type", "application/json")

			bearerToken := strings.TrimSpace(r.Header.Get("Authorization"))
			if bearerToken == "" {
				httputils.DispatchHTTPError(rw, "It was not possible to get the token from the headers", http.StatusForbidden)
				return
			}

			tokenParts := strings.Split(bearerToken, " ")
			if tokenParts[0] != "Bearer" {
				httputils.DispatchHTTPError(rw, "It should be Bearer authentication", http.StatusForbidden)
				return
			}

			user, err := token.GetUserFromToken(tokenParts[1])
			if err != nil {
				httpError := usererrormapper.MapUserErrorToHTTPError(err)
				httputils.DispatchHTTPError(rw, httpError.Message, httpError.StatusCode)
				return
			}

			newHTTPContext := context.WithValue(r.Context(), "user", user)
			hf(rw, r.WithContext(newHTTPContext))
		}
	}
}

// AdminOperation middleware ensure that only admins can execute a specific http handler
func AdminOperation() middleware.Middleware {
	return func(hf http.HandlerFunc) http.HandlerFunc {
		return func(rw http.ResponseWriter, r *http.Request) {
			rw.Header().Set("Content-type", "application/json")

			user, ok := r.Context().Value("user").(userdto.UserLoginDTO)
			if !ok {
				httputils.DispatchHTTPError(rw, "User should have been present", http.StatusForbidden)
				return
			}

			if !user.Admin {
				httputils.DispatchHTTPError(rw, "Only admins can perform this operation", http.StatusForbidden)
				return
			}

			hf(rw, r)
		}
	}
}

// EnsureEmailVerified .
func EnsureEmailVerified() middleware.Middleware {
	return func(hf http.HandlerFunc) http.HandlerFunc {
		return func(rw http.ResponseWriter, r *http.Request) {
			rw.Header().Set("Content-type", "application/json")

			user, ok := r.Context().Value("user").(userdto.UserLoginDTO)
			if !ok {
				httputils.DispatchHTTPError(rw, "User should have been present", http.StatusForbidden)
				return
			}

			if !user.EmailVerified {
				httputils.DispatchHTTPError(rw, "You must have your email verified", http.StatusForbidden)
				return
			}

			hf(rw, r)
		}
	}
}
