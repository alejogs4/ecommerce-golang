package productsdto

import "github.com/alejogs4/hn-website/src/products/domain/product"

type ProductDTO struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Picture     string  `json:"picture"`
	State       string  `json:"state"`
	Quantity    int     `json:"quantity"`
	Price       float64 `json:"price"`
	CreatedAt   string  `json:"created_at"`
}

func FromEntity(incommingProduct product.Product) ProductDTO {
	return ProductDTO{
		ID:          incommingProduct.ID(),
		Name:        incommingProduct.Name(),
		Description: incommingProduct.Description(),
		Picture:     incommingProduct.Picture(),
		State:       incommingProduct.State(),
		Price:       incommingProduct.Price(),
		Quantity:    incommingProduct.Quantity(),
		CreatedAt:   incommingProduct.CreatedAt().String(),
	}
}
