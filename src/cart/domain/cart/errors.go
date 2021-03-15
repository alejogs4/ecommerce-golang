package cart

import "errors"

// Errors with cart operations
var (
	ErrBadCartData                = errors.New("Cart: Cart fields must be provided")
	ErrUnauthorizedAction         = errors.New("Cart: You arent't authorized to make this action over the cart")
	ErrNotExistingCart            = errors.New("Cart: User has not any in progress cart")
	ErrInvalidCartStateTransition = errors.New("Cart: Invalid cart state transition")
	ErrInvalidCartState           = errors.New("Cart: Invalid cart state")
)
