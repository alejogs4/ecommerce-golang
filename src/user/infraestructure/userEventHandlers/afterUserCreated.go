package usereventhandlers

import (
	"errors"

	"github.com/alejogs4/hn-website/src/shared/domain/domainevent"
	mailservice "github.com/alejogs4/hn-website/src/shared/infraestructure/email/mailService"
	"github.com/alejogs4/hn-website/src/shared/infraestructure/token"
	userdto "github.com/alejogs4/hn-website/src/user/application/userDTO"
	"github.com/alejogs4/hn-website/src/user/domain/user"
)

// ErrEmailNoUserInformationFound error dispatched when in event information is not possible to retrieve user information
var ErrEmailNoUserInformationFound = errors.New("no user information was found in event")

// AfterUserCreatedWelcomeEmail event handler for send Welcome email
type AfterUserCreatedWelcomeEmail struct {
	MailService mailservice.Service
}

// Run method for send welcome email
func (aucwl AfterUserCreatedWelcomeEmail) Run(event domainevent.DomainEvent) error {

	userInfo, ok := event.EventInformation().(user.User)
	if !ok {
		return ErrEmailNoUserInformationFound
	}

	userDTO := userdto.FromRawUserToLoginUser(userInfo)
	mailInformation := mailservice.Mail{
		From:     "alejogs4@gmail.com",
		To:       []string{userDTO.Email},
		Subject:  "Welcome to our store, begin to use the best store house in the world!",
		Body:     userDTO,
		Template: "./mailTemplates/welcomeMail.html",
	}

	return aucwl.MailService.Send(mailInformation)
}

// AfterUserCreatedEmailVerification event handler for send verification email
type AfterUserCreatedEmailVerification struct {
	MailService mailservice.Service
}

// Run method for send verification email
func (aucev AfterUserCreatedEmailVerification) Run(event domainevent.DomainEvent) error {

	userInfo, ok := event.EventInformation().(user.User)
	if !ok {
		return ErrEmailNoUserInformationFound
	}

	userDTO := userdto.FromRawUserToLoginUser(userInfo)
	generatedToken, err := token.CreateToken(userDTO)
	if err != nil {
		return err
	}

	mailInformation := mailservice.Mail{
		From:    "alejogs4@gmail.com",
		To:      []string{userDTO.Email},
		Subject: "Verify your email to begin to use our services",
		Body: struct {
			userdto.UserLoginDTO
			Token string
		}{
			UserLoginDTO: userdto.UserLoginDTO{
				ID:            userDTO.ID,
				Name:          userDTO.Name,
				Lastname:      userDTO.Lastname,
				Email:         userDTO.Email,
				EmailVerified: userDTO.EmailVerified,
				Admin:         userDTO.Admin,
			},
			Token: generatedToken,
		},
		Template: "./mailTemplates/verificationMail.html",
	}

	return aucev.MailService.Send(mailInformation)
}
