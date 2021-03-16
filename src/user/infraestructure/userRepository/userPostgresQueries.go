package userrepository

import (
	"database/sql"
	"errors"

	"github.com/alejogs4/hn-website/src/user/domain/user"
)

type UserRepositoryPostgresQueries struct {
	db *sql.DB
}

func NewUserRepositoryPostgresQueries(db *sql.DB) UserRepositoryPostgresQueries {
	return UserRepositoryPostgresQueries{db: db}
}

func (ur UserRepositoryPostgresQueries) GetByID(id string) (user.User, error) {
	result := ur.db.QueryRow("SELECT id, name, email_verified, lastname, email, admin FROM users WHERE id=$1", id)

	var userID string
	var name string
	var emailVerified bool
	var lastname string
	var email string
	var admin bool

	err := result.Scan(&userID, &name, &emailVerified, &lastname, &email, &admin)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return user.User{}, user.ErrNotFoundUser
		}
		return user.User{}, err
	}

	gotUser := user.FromPrimitives(userID, name, lastname, email, admin, emailVerified)
	return gotUser, nil
}
