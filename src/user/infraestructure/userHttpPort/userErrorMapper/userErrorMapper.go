package usererrormapper

import (
	"errors"
	"net/http"

	httputils "github.com/alejogs4/hn-website/src/shared/infraestructure/httpUtils"
	"github.com/alejogs4/hn-website/src/user/domain/user"
)

// MapUserErrorToHTTPError .
func MapUserErrorToHTTPError(err error) httputils.HTTPError {
	switch {
	case errors.Is(err, user.ErrBadUserData):
		return httputils.HTTPError{Message: err.Error(), StatusCode: http.StatusBadRequest}
	case errors.Is(err, user.ErrInvalidUser):
		return httputils.HTTPError{Message: err.Error(), StatusCode: http.StatusBadRequest}
	case errors.Is(err, user.ErrInvalidAuth):
		return httputils.HTTPError{Message: err.Error(), StatusCode: http.StatusForbidden}
	case errors.Is(err, user.ErrTooShortPassword):
		return httputils.HTTPError{Message: err.Error(), StatusCode: http.StatusBadRequest}
	default:
		return httputils.HTTPError{Message: "Something went wrong", StatusCode: http.StatusInternalServerError}
	}
}
