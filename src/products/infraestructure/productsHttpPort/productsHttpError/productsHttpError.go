package productshttperror

import (
	"errors"
	"net/http"

	"github.com/alejogs4/hn-website/src/products/domain/product"
	httputils "github.com/alejogs4/hn-website/src/shared/infraestructure/httpUtils"
)

// MapProductErrorToHTTPError .
func MapProductErrorToHTTPError(err error) httputils.HTTPError {
	switch {
	case errors.Is(err, product.ErrBadProductData):
	case errors.Is(err, product.ErrZeroOrNegativeProductPrice):
	case errors.Is(err, product.ErrProductQuantity):
	case errors.Is(err, product.ErrNoEnoughQuantity):
	case errors.Is(err, product.ErrInvalidState):
		return httputils.HTTPError{Message: err.Error(), StatusCode: http.StatusBadRequest}
	case errors.Is(err, product.ErrNoExistingProduct):
		return httputils.HTTPError{Message: err.Error(), StatusCode: http.StatusNotFound}
	default:
		return httputils.HTTPError{Message: "Something went wrong", StatusCode: http.StatusInternalServerError}
	}

	return httputils.HTTPError{Message: "Something went wrong", StatusCode: http.StatusInternalServerError}
}
