package cart

// QueriesRepository to fetch information about user cart
type QueriesRepository interface {
	GetUserCart(userID string) (Cart, error)
	ExistUserCart(userID string) (bool, error)
	GetCartItemCount(productID, cartID string) (int, error)
	GetCartProductsCount(cartID string) (int, error)
}

// CommandsRepository to add remove and finally buy user cart items
type CommandsRepository interface {
	CreateCart(cart Cart) error
	AddCartItem(item Item, cartID string) error
	RemoveCartItem(cartItemID string) error
	RemoveCart(cartID string) error
	BuyCart(cartID string) error
}
