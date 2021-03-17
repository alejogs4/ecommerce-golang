package productevents

import (
	"errors"

	cartdto "github.com/alejogs4/hn-website/src/cart/application/cartDTO"
	"github.com/alejogs4/hn-website/src/cart/domain/cart"
	"github.com/alejogs4/hn-website/src/products/domain/product"
	"github.com/alejogs4/hn-website/src/shared/domain/domainevent"
)

type AfterCarWasOrdered struct {
	ProductCommands product.CommandsRepository
}

func (ac AfterCarWasOrdered) Run(event domainevent.DomainEvent) error {
	boughtCart, ok := event.EventInformation().(cart.Cart)
	if !ok {
		return errors.New("Cart: cart was not found")
	}

	boughtProducts := cartdto.ProductQuantityFromCartItems(boughtCart.GetProducts())
	for _, boughtProduct := range boughtProducts {
		err := ac.ProductCommands.BuyProduct(boughtProduct.BoughtProduct, boughtProduct.BoughtQuantity)
		if err != nil {
			return err
		}
	}

	return nil
}
