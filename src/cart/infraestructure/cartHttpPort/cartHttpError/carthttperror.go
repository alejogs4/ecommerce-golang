package carthttperror

import (
	"errors"
	"net/http"

	"github.com/alejogs4/hn-website/src/cart/domain/cart"
	httputils "github.com/alejogs4/hn-website/src/shared/infraestructure/httpUtils"
)

// MapCartErrorToHTTPError .
func MapCartErrorToHTTPError(err error) httputils.HTTPError {
	switch {
	case errors.Is(err, cart.ErrBadCartData):
	case errors.Is(err, cart.ErrInvalidCartStateTransition):
	case errors.Is(err, cart.ErrInvalidCartState):
	case errors.Is(err, cart.ErrInvalidCartItemState):
		return httputils.HTTPError{Message: err.Error(), StatusCode: http.StatusBadRequest}
	case errors.Is(err, cart.ErrUnauthorizedAction):
		return httputils.HTTPError{Message: err.Error(), StatusCode: http.StatusForbidden}
	case errors.Is(err, cart.ErrNotExistingCart):
		return httputils.HTTPError{Message: err.Error(), StatusCode: http.StatusNotFound}
	}

	return httputils.HTTPError{Message: "Something went wrong", StatusCode: http.StatusInternalServerError}
}
