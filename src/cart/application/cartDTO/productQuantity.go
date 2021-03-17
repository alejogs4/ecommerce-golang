package cartdto

import (
	"github.com/alejogs4/hn-website/src/cart/domain/cart"
	"github.com/alejogs4/hn-website/src/products/domain/product"
)

// ProductWithQuantity struct to hold besides the product the quantity has been bought
type ProductWithQuantity struct {
	BoughtProduct  product.Product
	BoughtQuantity int
}

// ProductQuantityFromCartItems get from cart items how many of its individual products has been bought by the user
func ProductQuantityFromCartItems(items []cart.Item) map[string]ProductWithQuantity {
	var products []product.Product
	for _, cartItem := range items {
		products = append(products, cartItem.Product())
	}

	return groupProductsByID(products)
}

func groupProductsByID(products []product.Product) map[string]ProductWithQuantity {
	productsByID := make(map[string]ProductWithQuantity)

	for _, product := range products {
		accumulatedProduct := productsByID[product.ID()]

		productsByID[product.ID()] = ProductWithQuantity{
			BoughtProduct:  product,
			BoughtQuantity: accumulatedProduct.BoughtQuantity + 1,
		}
	}

	return productsByID
}
