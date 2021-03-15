package token

import (
	"errors"
	"time"

	userdto "github.com/alejogs4/hn-website/src/user/application/userDTO"
	"github.com/alejogs4/hn-website/src/user/domain/user"
	"github.com/dgrijalva/jwt-go"
)

type claimer struct {
	userdto.UserLoginDTO
	jwt.StandardClaims
}

// CreateToken create new jwt string based in the user passed as payload, besides returns a error if something went wrong generating the token
func CreateToken(payload userdto.UserLoginDTO) (string, error) {
	tokenClaimer := claimer{
		UserLoginDTO: payload,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
			Issuer:    "store",
		},
	}

	signedToken := jwt.NewWithClaims(jwt.SigningMethodRS256, tokenClaimer)
	return signedToken.SignedString(signingKey)
}

// GetUserFromToken returns a user from the token or error if this one is invalid or another problem
func GetUserFromToken(token string) (userdto.UserLoginDTO, error) {
	parsedToken, err := jwt.ParseWithClaims(token, &claimer{}, verifyFunction)
	if err != nil {
		return userdto.UserLoginDTO{}, err
	}

	if !parsedToken.Valid {
		return userdto.UserLoginDTO{}, user.ErrInvalidAuth
	}

	claim, ok := parsedToken.Claims.(*claimer)
	if !ok {
		return userdto.UserLoginDTO{}, errors.New("It was not possible get user from token")
	}

	return claim.UserLoginDTO, nil
}

func verifyFunction(t *jwt.Token) (interface{}, error) {
	return verifyKey, nil
}
