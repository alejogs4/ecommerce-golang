package product

import "errors"

// Errors for products
var (
	ErrBadProductData             = errors.New("Product: Product must be provided with all its fields")
	ErrZeroOrNegativeProductPrice = errors.New("Product: Product price must have a positive value")
	ErrProductQuantity            = errors.New("Product: Product quantity cannot be negative")
	ErrNoEnoughQuantity           = errors.New("Product: There are not enought products")
	ErrInvalidState               = errors.New("Product: invalid product state")
)
