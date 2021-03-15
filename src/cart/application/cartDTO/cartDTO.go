package cartdto

import (
	"github.com/alejogs4/hn-website/src/cart/domain/cart"
	productsdto "github.com/alejogs4/hn-website/src/products/application/productsDTO"
)

type CartDTO struct {
	ID        string        `json:"id"`
	UserID    string        `json:"user_id"`
	State     string        `json:"state"`
	CreatedAt string        `json:"created_at"`
	Products  []CartItemDTO `json:"items"`
}

type CartItemDTO struct {
	ID      string                 `json:"id"`
	state   string                 `json:"state"`
	Product productsdto.ProductDTO `json:"product"`
}

func FromEntity(incommingCart cart.Cart) CartDTO {
	var items []CartItemDTO = []CartItemDTO{}
	for _, product := range incommingCart.GetProducts() {
		items = append(items, CartItemDTO{
			ID:      product.ID(),
			state:   product.State(),
			Product: productsdto.FromEntity(product.Product()),
		})
	}

	return CartDTO{
		ID:        incommingCart.GetID(),
		UserID:    incommingCart.GetUserID(),
		State:     incommingCart.GetState(),
		CreatedAt: incommingCart.GetCreatedAt().String(),
		Products:  items,
	}
}
