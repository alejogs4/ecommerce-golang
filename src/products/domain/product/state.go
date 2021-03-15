package product

// State string which will represent the current state of a product
type State string

const (
	// Active state of a product
	Active State = "ACTIVE"
	// UnAvailable state of a product
	UnAvailable State = "UNAVAILABLE"
	// Removed state of a product
	Removed State = "REMOVED"
)

func NewState(str string) (State, error) {
	if str != Active.String() && str != UnAvailable.String() && str != Removed.String() {
		return State(""), ErrInvalidState
	}

	return State(str), nil
}

func (ps State) String() string {
	return string(ps)
}
