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

func (ps State) String() string {
	return string(ps)
}
