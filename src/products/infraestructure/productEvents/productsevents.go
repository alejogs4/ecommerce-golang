package productevents

import (
	"errors"

	"github.com/alejogs4/hn-website/src/cart/domain/cart"
	"github.com/alejogs4/hn-website/src/products/domain/product"
	"github.com/alejogs4/hn-website/src/shared/domain/domainevent"
)

type AfterCarWasOrdered struct {
	ProductCommands product.CommandsRepository
}

type ProductWithQuantity struct {
	BoughtProduct  product.Product
	BoughtQuantity int
}

func (ac AfterCarWasOrdered) Run(event domainevent.DomainEvent) error {
	boughtCart, ok := event.EventInformation().(cart.Cart)
	if !ok {
		return errors.New("Cart: cart was not found")
	}

	var products []product.Product
	for _, cartItem := range boughtCart.GetProducts() {
		products = append(products, cartItem.Product())
	}

	boughtProducts := ac.groupProductsByID(products)
	for _, boughtProduct := range boughtProducts {
		err := ac.ProductCommands.BuyProduct(boughtProduct.BoughtProduct, boughtProduct.BoughtQuantity)
		if err != nil {
			return err
		}
	}

	return nil
}

func (ac AfterCarWasOrdered) groupProductsByID(products []product.Product) map[string]ProductWithQuantity {
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
