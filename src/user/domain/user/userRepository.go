package user

// Queries for users information
type Queries interface {
	GetByID(id string) (User, error)
}

// CommandsRepository for manipulate user information and their authetication and authorization in the system
type CommandsRepository interface {
	CreateUser(user User) error
	LoginUser(email string, password string) (User, error)
	VerifyEmail(userEmail string) error
}
