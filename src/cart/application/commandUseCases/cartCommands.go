package commandusecases

import (
	"github.com/alejogs4/hn-website/src/cart/domain/cart"
	"github.com/alejogs4/hn-website/src/products/domain/product"
	"github.com/google/uuid"
)

// CartCommands struct which will handle the application rules for add new cart item
type CartCommands struct {
	commands       cart.CommandsRepository
	queries        cart.QueriesRepository
	productQueries product.QueriesRepository
}

// NewCartCommand returns a new instance of CartCommands
func NewCartCommand(commands cart.CommandsRepository, queries cart.QueriesRepository, productQueries product.QueriesRepository) CartCommands {
	return CartCommands{commands: commands, queries: queries, productQueries: productQueries}
}

// AddCartItem validate if new cart is necessary for the user and at the end add the item if rules apply
func (cc *CartCommands) AddCartItem(userID, productID string) (string, error) {
	existCart, err := cc.queries.ExistUserCart(userID)
	if err != nil {
		return "", err
	}

	userCart := cart.Cart{}
	cartID := ""

	// Create cart if doesn't exist
	if !existCart {
		cartID = uuid.New().String()

		userCart, err = cart.NewCart(cartID, userID)
		if err != nil {
			return "", err
		}

		err = cc.commands.CreateCart(userCart)
		if err != nil {
			return "", err
		}
	}

	// If the cart exist fetch it from database otherwise, continues with new just created one
	if cartID == "" {
		userCart, err = cc.queries.GetUserCart(userID)
		if err != nil {
			return "", err
		}
		cartID = userCart.GetID()
	}

	cartProduct, err := cc.productQueries.GetByID(productID)
	if err != nil {
		return "", err
	}

	productCartCount, err := cc.queries.GetCartItemCount(productID, cartID)
	if err != nil {
		return "", err
	}

	// If the fact of add this item to the cart exceeds the current ones, return an error
	if productCartCount+1 > cartProduct.Quantity() {
		return "", product.ErrNoEnoughQuantity
	}

	// TODO: Reduce product quantity by one

	cartItemID := uuid.New().String()
	cartItem := cart.NewCartItem(cartItemID, cartProduct, cart.ADDED)
	if err := cc.commands.AddCartItem(cartItem, cartID); err != nil {
		return "", err
	}

	return cartItemID, nil
}

// RemoveCartItem remove cart item and cart itself if the removed cart item was the last one remaining item
func (cc *CartCommands) RemoveCartItem(userID, itemID, cartID string) error {
	currentCart, err := cc.queries.GetUserCart(userID)
	if err != nil {
		return err
	}

	if currentCart.GetID() != cartID {
		return cart.ErrUnauthorizedAction
	}

	err = cc.commands.RemoveCartItem(itemID)
	if err != nil {
		return err
	}

	productCount, err := cc.queries.GetCartProductsCount(cartID)
	if err != nil {
		return err
	}

	if productCount <= 0 {
		return cc.commands.RemoveCart(cartID)
	}

	return nil
}
