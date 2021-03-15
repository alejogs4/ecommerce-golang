package userrepository

import (
	"database/sql"
	"errors"

	"github.com/alejogs4/hn-website/src/user/domain/user"
	"golang.org/x/crypto/bcrypt"
)

// UserPostgresCommandsRepository concrete implementation of user Commands repository
type UserPostgresCommandsRepository struct {
	db *sql.DB
}

// NewUserPostgresCommandsRepository returns a new instance of UserPostgresCommands
func NewUserPostgresCommandsRepository(db *sql.DB) UserPostgresCommandsRepository {
	return UserPostgresCommandsRepository{db: db}
}

// CreateUser insert new user in postgres database
func (upc UserPostgresCommandsRepository) CreateUser(newUser user.User) error {
	_, err := upc.db.Exec(`
		INSERT INTO users(id, name, lastname, email, password, admin, email_verified)
		VALUES($1, $2, $3, $4, $5, $6, $7)
	`,
		newUser.GetID(),
		newUser.GetName(),
		newUser.GetLastname(),
		newUser.GetEmail(),
		newUser.GetPassword(),
		newUser.IsAdmin(),
		newUser.HasEmailVerified())

	return err
}

// LoginUser verify is passed email and password are from an existing user in the database
func (upc UserPostgresCommandsRepository) LoginUser(email string, password string) (user.User, error) {
	row := upc.db.QueryRow(`
	SELECT id, name, lastname, admin, email_verified, password FROM users WHERE email=$1
	`, email)

	var id string
	var name string
	var lastname string
	var userPassword string
	var admin bool
	var emailVerified bool
	err := row.Scan(&id, &name, &lastname, &admin, &emailVerified, &userPassword)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return user.User{}, user.ErrInvalidUser
		}
		return user.User{}, err
	}

	returnedUser, err := user.NewUser(id, name, lastname, email, password, admin, emailVerified)
	if err != nil {
		return user.User{}, err
	}
	// This could be done inside application code, since it shouldn't depends which repository implements it
	err = bcrypt.CompareHashAndPassword([]byte(userPassword), []byte(password))
	if err != nil {
		return user.User{}, user.ErrInvalidUser
	}

	return returnedUser, nil
}

// VerifyEmail update user email verify, this is because user was sent a verification email to his/her email
// this way we avoid unfair email check at application level since some emails can be seen tricky to verify by itself
func (upc UserPostgresCommandsRepository) VerifyEmail(userEmail string) error {
	_, err := upc.db.Exec("UPDATE users SET email_verified=$1 WHERE email=$2", true, userEmail)
	return err
}
