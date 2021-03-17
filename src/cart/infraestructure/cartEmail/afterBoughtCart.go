package cartemail

import (
	"errors"
	"fmt"

	cartdto "github.com/alejogs4/hn-website/src/cart/application/cartDTO"
	"github.com/alejogs4/hn-website/src/cart/domain/cart"
	productsdto "github.com/alejogs4/hn-website/src/products/application/productsDTO"
	"github.com/alejogs4/hn-website/src/shared/domain/domainevent"
	mailservice "github.com/alejogs4/hn-website/src/shared/infraestructure/email/mailService"
	"github.com/alejogs4/hn-website/src/user/domain/user"
)

// ErrNoFoundCart is dispacthed when a car was not sent to the email handler
var ErrNoFoundCart = errors.New("Cart: No cart received in the event")

// AfterBoughtCart email handler
type AfterBoughtCart struct {
	EmailService mailservice.Service
	UserQueries  user.Queries
}

type ProductDTOQuantity struct {
	Product  productsdto.ProductDTO
	Quantity int
}

// Run to send an email to notify the user that its products has been succesfully bought
func (abc AfterBoughtCart) Run(event domainevent.DomainEvent) error {
	boughtCart, ok := event.EventInformation().(cart.Cart)
	if !ok {
		return ErrNoFoundCart
	}

	cartUser, err := abc.UserQueries.GetByID(boughtCart.GetUserID())
	if err != nil {
		return err
	}

	var cartDTOWithQuantity []ProductDTOQuantity
	productsBoughtQuantity := cartdto.ProductQuantityFromCartItems(boughtCart.GetProducts())

	for _, cartItemInfo := range productsBoughtQuantity {
		cartDTOWithQuantity = append(cartDTOWithQuantity, ProductDTOQuantity{
			Quantity: cartItemInfo.BoughtQuantity,
			Product:  productsdto.FromEntity(cartItemInfo.BoughtProduct),
		})
	}

	boughtCartEmail := mailservice.Mail{
		From:     "alejogs4@gmail.com",
		To:       []string{cartUser.GetEmail()},
		Subject:  fmt.Sprintf("Cart was bought succesfully with %d products", len(boughtCart.GetProducts())),
		Template: "./mailTemplates/cart-bougth.html",
		Body:     cartDTOWithQuantity,
	}

	return abc.EmailService.Send(boughtCartEmail)
}
