package userhttpport

import (
	mailservice "github.com/alejogs4/hn-website/src/shared/infraestructure/email/mailService"
	usercommands "github.com/alejogs4/hn-website/src/user/application/userCommands"
	"github.com/alejogs4/hn-website/src/user/domain/user"
	userevents "github.com/alejogs4/hn-website/src/user/domain/user/userEvents"
	usereventhandlers "github.com/alejogs4/hn-website/src/user/infraestructure/userEventHandlers"
	"github.com/gorilla/mux"
)

// HandleUserControllers register all routes regarding user
func HandleUserControllers(router *mux.Router, userCommandsRepository user.CommandsRepository, emailService mailservice.Service) {
	commands := usercommands.NewUserCommandsUseCases(userCommandsRepository)
	contollers := commandsControllers{UseCases: commands}

	// Register event handlers
	commands.RegisterEventHandler(userevents.UserCreatedEvent, usereventhandlers.AfterUserCreatedWelcomeEmail{MailService: emailService})
	commands.RegisterEventHandler(userevents.UserCreatedEvent, usereventhandlers.AfterUserCreatedEmailVerification{MailService: emailService})

	router.HandleFunc("/api/v1/signup", contollers.CreateUserController)
	router.HandleFunc("/api/v1/login", contollers.LoginUserController)
	router.HandleFunc("/api/v1/user/verify/email", contollers.VerifyUserEmail)
}
