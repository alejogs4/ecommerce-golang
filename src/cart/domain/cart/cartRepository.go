package cart

// Queries to fetch information about user cart
type Queries interface {
	GetUserCart(userID string) (Cart, error)
	ExistUserCart(userID string) bool
	GetCartItemCount(productID, cartID string) (int, error)
	GetCartProductsCount(cartID string) (int, error)
}

// Commands to add remove and finally buy user cart items
type Commands interface {
	CreateCart(cart Cart) error
	AddCartItem(item Item, cartID string) error
	RemoveCartItem(cartItemID, userID string) error
	RemoveCart(cartID string) error
	BuyCart(cartID string) error
}
