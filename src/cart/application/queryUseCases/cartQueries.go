package queryusecases

import (
	cartdto "github.com/alejogs4/hn-website/src/cart/application/cartDTO"
	"github.com/alejogs4/hn-website/src/cart/domain/cart"
)

// CartQueries use cases to fetch user information
type CartQueries struct {
	queries cart.QueriesRepository
}

// NewCartQueries returns a new instance of CartQueries use cases
func NewCartQueries(queries cart.QueriesRepository) CartQueries {
	return CartQueries{queries: queries}
}

// GetUserCart validates if inProgress cart exists beforehand besides verify user fetching cart is the same which is the owner
func (cq *CartQueries) GetUserCart(userID string) (cartdto.CartDTO, error) {
	existCart, err := cq.queries.ExistUserCart(userID)

	if err != nil {
		return cartdto.CartDTO{}, err
	}

	if !existCart {
		return cartdto.CartDTO{}, cart.ErrNotExistingCart
	}

	userCart, err := cq.queries.GetUserCart(userID)
	if err != nil {
		return cartdto.CartDTO{}, err
	}

	return cartdto.FromEntity(userCart), nil
}
