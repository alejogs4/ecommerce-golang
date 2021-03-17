package user

import (
	"unicode/utf8"

	"github.com/alejogs4/hn-website/src/shared/domain/aggregate"
	"github.com/alejogs4/hn-website/src/shared/domain/valueobject"
	userevents "github.com/alejogs4/hn-website/src/user/domain/user/userEvents"
)

const minPasswordLength = 6

// User hold user information
type User struct {
	id            valueobject.MaybeEmpty
	name          valueobject.MaybeEmpty
	emailVerified bool
	lastname      valueobject.MaybeEmpty
	email         valueobject.MaybeEmpty
	password      valueobject.MaybeEmpty
	admin         bool
	aggregate.CommonAggregate
}

// NewUser returns a new user entitiy validating passed fields
func NewUser(id, name, lastname, email, password string, admin, emailVerified bool) (User, error) {
	validID := valueobject.NewMaybeEmpty(id)
	validName := valueobject.NewMaybeEmpty(name)
	validLastname := valueobject.NewMaybeEmpty(lastname)
	validEmail := valueobject.NewMaybeEmpty(email)
	validPassword := valueobject.NewMaybeEmpty(password)

	if validID.IsEmpty() || validName.IsEmpty() || validLastname.IsEmpty() || validEmail.IsEmpty() || validPassword.IsEmpty() {
		return User{}, ErrBadUserData
	}

	if utf8.RuneCount([]byte(validPassword.String())) < minPasswordLength {
		return User{}, ErrTooShortPassword
	}
	createdUser := User{
		id:            validID,
		name:          validName,
		lastname:      validLastname,
		email:         validEmail,
		password:      validPassword,
		admin:         admin,
		emailVerified: emailVerified,
	}

	createdUser.RegisterEvent(userevents.UserCreated{Information: createdUser})

	return createdUser, nil
}

// FromPrimitives returns a user from its most primitive values, this function assumes that its user will pass right values
func FromPrimitives(id, name, lastname, email string, admin, emailVerified bool) User {
	return User{
		id:            valueobject.NewMaybeEmpty(id),
		name:          valueobject.NewMaybeEmpty(name),
		emailVerified: emailVerified,
		lastname:      valueobject.NewMaybeEmpty(lastname),
		email:         valueobject.NewMaybeEmpty(email),
		password:      "",
		admin:         admin,
	}
}

// GetID .
func (u *User) GetID() string {
	return u.id.String()
}

// GetName .
func (u *User) GetName() string {
	return u.name.String()
}

// GetLastname .
func (u *User) GetLastname() string {
	return u.lastname.String()
}

// GetEmail .
func (u *User) GetEmail() string {
	return u.email.String()
}

// GetPassword .
func (u *User) GetPassword() string {
	return u.password.String()
}

// IsAdmin .
func (u *User) IsAdmin() bool {
	return u.admin
}

// HasEmailVerified .
func (u *User) HasEmailVerified() bool {
	return u.emailVerified
}
