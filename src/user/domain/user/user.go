package user

import (
	"log"
	"strings"
	"unicode/utf8"

	"github.com/alejogs4/hn-website/src/shared/domain/domainevent"
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
	events        []domainevent.DomainEvent
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

	createdUser.registerEvent(userevents.UserCreated{Information: createdUser})

	return createdUser, nil
}

func (u *User) registerEvent(event domainevent.DomainEvent) {
	u.events = append(u.events, event)
}

// DispatchRegisteredEvents implementation to execute user events handlers
func (u *User) DispatchRegisteredEvents(eventHandlers map[string][]domainevent.DomainEventHandler, targetEvents []string) {
	// TODO: Look how abstract this in a separate struct in order to reduce repetition, this will be the same in all
	eventsString := strings.Join(targetEvents, " ")

	for _, e := range u.events {
		if !strings.Contains(eventsString, e.Name()) {
			continue
		}

		if eventHandlers, ok := eventHandlers[e.Name()]; ok {
			for _, handler := range eventHandlers {
				go func(hn domainevent.DomainEventHandler, event domainevent.DomainEvent) {
					err := hn.Run(event)
					if err != nil {
						// It's a decision that event handler errors will not stop application
						log.Printf("Error: %s in event: %v", err, event)
					}
				}(handler, e)
			}
		}
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
