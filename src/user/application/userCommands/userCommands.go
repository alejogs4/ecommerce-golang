package usercommands

import (
	usecase "github.com/alejogs4/hn-website/src/shared/domain/useCase"
	userdto "github.com/alejogs4/hn-website/src/user/application/userDTO"
	"github.com/alejogs4/hn-website/src/user/domain/user"
	userevents "github.com/alejogs4/hn-website/src/user/domain/user/userEvents"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

// UseCases uses cases for application user
type UseCases struct {
	commands user.CommandsRepository
	usecase.EventScheduler
}

// NewUserCommandsUseCases returns a new instance of UserCommandsUseCases
func NewUserCommandsUseCases(commands user.CommandsRepository) UseCases {
	return UseCases{commands: commands, EventScheduler: usecase.NewEventScheduler()}
}

// LoginUser executes and verify user login and before it hash incoming password with a cost of 14
func (uc *UseCases) LoginUser(email, password string) (userdto.UserLoginDTO, error) {
	rawUser, err := uc.commands.LoginUser(email, password)
	if err != nil {
		return userdto.UserLoginDTO{}, err
	}

	return userdto.FromRawUserToLoginUser(rawUser), nil
}

// CreateUser verify user incomming information and hash the password to create it at the end, after these validations
func (uc *UseCases) CreateUser(name, lastname, email, password string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return err
	}

	createdUser, err := user.NewUser(uuid.New().String(), name, lastname, email, string(hashedPassword), false, false)
	if err != nil {
		return err
	}

	go createdUser.DispatchRegisteredEvents(uc.Handlers(), []string{userevents.UserCreatedEvent})
	return uc.commands.CreateUser(createdUser)
}

// VerifyEmail executes user email verify in order to approve the use of user registered email
func (uc *UseCases) VerifyEmail(userEmail string) error {
	return uc.commands.VerifyEmail(userEmail)
}
