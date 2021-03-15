package cart

import (
	"time"

	"github.com/alejogs4/hn-website/src/shared/domain/valueobject"
)

// Cart struct which represent an user list of products to be bougth
type Cart struct {
	id        valueobject.MaybeEmpty
	userID    valueobject.MaybeEmpty
	state     State
	createdAt time.Time
	products  []Item
}

// NewCart check provided data and validates that is not in an invalid state
func NewCart(id, userID string) (Cart, error) {
	cartID := valueobject.NewMaybeEmpty(id)
	cartUserID := valueobject.NewMaybeEmpty(userID)

	if cartID.IsEmpty() || cartUserID.IsEmpty() {
		return Cart{}, ErrBadCartData
	}

	return Cart{
		id:        cartID,
		userID:    cartUserID,
		products:  []Item{},
		state:     InProgress,
		createdAt: time.Now(),
	}, nil
}

// FromPrimitives returns a new cart from its most primitive values as plain strings
func FromPrimitives(id, userID, state, createdAt string, products []Item) (Cart, error) {
	cartID := valueobject.NewMaybeEmpty(id)
	cartUserID := valueobject.NewMaybeEmpty(userID)
	cartState := valueobject.NewMaybeEmpty(state)
	cartCreationDate := valueobject.NewMaybeEmpty(createdAt)

	if cartID.IsEmpty() || cartUserID.IsEmpty() || cartState.IsEmpty() || cartCreationDate.IsEmpty() {
		return Cart{}, ErrBadCartData
	}

	newState, err := NewState(cartState.String())
	if err != nil {
		return Cart{}, err
	}

	createdAtTime, err := time.Parse(time.RFC3339, createdAt)
	if err != nil {
		return Cart{}, err
	}

	return Cart{
		id:        cartID,
		userID:    cartUserID,
		state:     newState,
		products:  products,
		createdAt: createdAtTime,
	}, nil
}

func (c *Cart) GetID() string {
	return c.id.String()
}

func (c *Cart) GetUserID() string {
	return c.userID.String()
}

func (c *Cart) GetState() string {
	return string(c.state)
}

func (c *Cart) GetCreatedAt() time.Time {
	return c.createdAt
}

func (c *Cart) GetProducts() []Item {
	return c.products
}
