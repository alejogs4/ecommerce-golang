package cartevents

import "time"

// Bougth event will be dispacthed when a car has been bought for the user
const Bougth = "Cart_Bougth"

// CartBought event struct
type CartBought struct {
	Information interface{}
}

// Name .
func (cb CartBought) Name() string {
	return Bougth
}

// OccurredTimed .
func (cb CartBought) OccurredTimed() time.Time {
	return time.Now()
}

// EventInformation .
func (cb CartBought) EventInformation() interface{} {
	return cb.Information
}
