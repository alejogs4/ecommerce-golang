package queryusecases

import "github.com/alejogs4/hn-website/src/cart/domain/cart"

// CartQueries use cases to fetch user information
type CartQueries struct {
	queries cart.Queries
}

// NewCartQueries returns a new instance of CartQueries use cases
func NewCartQueries(queries cart.Queries) CartQueries {
	return CartQueries{queries: queries}
}

// GetUserCart validates if inProgress cart exists beforehand besides verify user fetching cart is the same which is the owner
func (cq *CartQueries) GetUserCart(userID string) (cart.Cart, error) {
	existCart := cq.queries.ExistUserCart(userID)
	if !existCart {
		return cart.Cart{}, cart.ErrNotExistingCart
	}

	return cq.queries.GetUserCart(userID)
}
