package cart

// State of the cart depending of how user interact with the system
type State string

const (
	// InProgress initial state of the cart, this is automatically set as soon as user add his/her first item to the cart
	InProgress State = "InProgress"
	// Ordered state set to the cart when user makes final purchase, this is a final state
	Ordered State = "Ordered"
	// Removed State is set when all items of an existing car (one previosly in InProgress state) are removed from it, this is a final state
	Removed State = "Removed"
)

// SetNewState acts as a simple state transitioning function, to set a new state as long as this is a valid one
func (s State) SetNewState(newState State) (State, error) {
	if s == Ordered || s == Removed {
		return s, ErrInvalidCartStateTransition
	}

	if newState != InProgress && newState != Ordered && newState != Removed {
		return s, ErrInvalidCartState
	}

	return newState, nil
}
