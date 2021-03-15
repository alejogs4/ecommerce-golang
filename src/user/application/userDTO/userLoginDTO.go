package userdto

import "github.com/alejogs4/hn-website/src/user/domain/user"

// UserLoginDTO structure to be returned to user in user login
type UserLoginDTO struct {
	ID            string `json:"id"`
	Name          string `json:"name"`
	EmailVerified bool   `json:"email_verified"`
	Lastname      string `json:"lastname"`
	Email         string `json:"email"`
	Admin         bool   `json:"admin"`
}

// FromRawUserToLoginUser receives a normal domain user and return the specific one to be return during the login
func FromRawUserToLoginUser(user user.User) UserLoginDTO {
	return UserLoginDTO{
		ID:            user.GetID(),
		Name:          user.GetName(),
		Lastname:      user.GetLastname(),
		Email:         user.GetEmail(),
		EmailVerified: user.HasEmailVerified(),
		Admin:         user.IsAdmin(),
	}
}
