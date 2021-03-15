package cart

import "github.com/alejogs4/hn-website/src/shared/domain/valueobject"

// Cart struct which represent an user list of products to be bougth
type Cart struct {
	id       valueobject.MaybeEmpty
	userID   valueobject.MaybeEmpty
	state    State
	products []Item
}

// NewCart check provided data and validates that is not in an invalid state
func NewCart(id, userID string) (Cart, error) {
	cartID := valueobject.NewMaybeEmpty(id)
	cartUserID := valueobject.NewMaybeEmpty(userID)

	if cartID.IsEmpty() || cartUserID.IsEmpty() {
		return Cart{}, ErrBadCartData
	}

	return Cart{
		id:       cartID,
		userID:   cartUserID,
		state:    InProgress,
		products: []Item{},
	}, nil
}

func (c *Cart) GetID() string {
	return c.id.String()
}
