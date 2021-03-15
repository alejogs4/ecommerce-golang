package userevents

import "time"

const UserCreatedEvent = "USER_CREATED"

// UserCreated event .
type UserCreated struct {
	Information interface{}
}

func (uc UserCreated) Name() string {
	return UserCreatedEvent
}

func (uc UserCreated) OccurredTimed() time.Time {
	return time.Now()
}

func (uc UserCreated) EventInformation() interface{} {
	return uc.Information
}
