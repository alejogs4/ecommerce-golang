package user

import (
	"errors"
	"fmt"
)

// Domain error for users
var (
	ErrBadUserData      = errors.New("User: All user fields must be present")
	ErrInvalidUser      = errors.New("User: user doesn't exist")
	ErrInvalidAuth      = errors.New("User: user failed authentication")
	ErrTooShortPassword = fmt.Errorf("User: password must be larger or equal to %d", minPasswordLength)
)
