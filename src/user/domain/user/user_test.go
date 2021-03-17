package user_test

import (
	"errors"
	"testing"

	"github.com/alejogs4/hn-website/src/user/domain/user"
	"github.com/icrowley/fake"
)

func TestNewUserEntity(t *testing.T) {
	rightUserMap := map[string]string{
		"id":       fake.IPv4(),
		"name":     fake.FirstName(),
		"lastname": fake.LastName(),
		"email":    fake.EmailAddress(),
		"password": fake.Password(7, 14, true, true, true),
	}

	testCases := []struct {
		Name          string
		StrFields     map[string]string
		ExpectedError error
		ExpectedUser  user.User
	}{
		{
			Name: "Should throw error if id is empty",
			StrFields: map[string]string{
				"name":     fake.FirstName(),
				"lastname": fake.LastName(),
				"email":    fake.EmailAddress(),
				"password": fake.Password(6, 10, true, true, true),
			},
			ExpectedError: user.ErrBadUserData,
			ExpectedUser:  user.User{},
		},
		{
			Name: "Should throw error if name is empty",
			StrFields: map[string]string{
				"id":       fake.IPv4(),
				"lastname": fake.LastName(),
				"email":    fake.EmailAddress(),
				"password": fake.Password(6, 10, true, true, true),
			},
			ExpectedError: user.ErrBadUserData,
			ExpectedUser:  user.User{},
		},
		{
			Name: "Should throw error if lastname is empty",
			StrFields: map[string]string{
				"id":       fake.IPv4(),
				"name":     fake.FirstName(),
				"email":    fake.EmailAddress(),
				"password": fake.Password(6, 10, true, true, true),
			},
			ExpectedError: user.ErrBadUserData,
			ExpectedUser:  user.User{},
		},
		{
			Name: "Should throw error if email is empty",
			StrFields: map[string]string{
				"id":       fake.IPv4(),
				"name":     fake.FirstName(),
				"lastname": fake.LastName(),
				"password": fake.Password(6, 10, true, true, true),
			},
			ExpectedError: user.ErrBadUserData,
			ExpectedUser:  user.User{},
		},
		{
			Name: "Should throw error if password is empty",
			StrFields: map[string]string{
				"id":       fake.IPv4(),
				"name":     fake.FirstName(),
				"lastname": fake.LastName(),
				"email":    fake.EmailAddress(),
			},
			ExpectedError: user.ErrBadUserData,
			ExpectedUser:  user.User{},
		},
		{
			Name: "Should throw error if password is too short",
			StrFields: map[string]string{
				"id":       fake.IPv4(),
				"name":     fake.FirstName(),
				"lastname": fake.LastName(),
				"email":    fake.EmailAddress(),
				"password": fake.Password(2, 4, true, true, true),
			},
			ExpectedError: user.ErrTooShortPassword,
			ExpectedUser:  user.User{},
		},
		{
			Name:          "Should return user if every field is right",
			StrFields:     rightUserMap,
			ExpectedError: nil,
			ExpectedUser:  user.FromPrimitives(rightUserMap["id"], rightUserMap["name"], rightUserMap["lastname"], rightUserMap["email"], false, false),
		},
	}

	for _, tCase := range testCases {
		t.Run(tCase.Name, func(t *testing.T) {
			gotUser, err := user.NewUser(tCase.StrFields["id"], tCase.StrFields["name"], tCase.StrFields["lastname"], tCase.StrFields["email"], tCase.StrFields["password"], false, false)

			if !errors.Is(err, tCase.ExpectedError) {
				t.Fatalf("Error: %v was expected, %v was got", tCase.ExpectedError, err)
			}
			if err == nil {
				if gotUser.GetID() != tCase.StrFields["id"] {
					t.Fatalf("Error: %s was expected, %s was got", tCase.StrFields["id"], gotUser.GetID())
				}
			}
		})
	}
}

func TestUserGetters(t *testing.T) {
	t.Run("Should return user id", func(t *testing.T) {
		expectedID := fake.IPv4()
		tUser := user.FromPrimitives(expectedID, "", "", "", false, false)

		if tUser.GetID() != expectedID {
			t.Fatalf("Error: %s was expected, %s was got", expectedID, tUser.GetID())
		}
	})

	t.Run("Should return user name", func(t *testing.T) {
		expectedName := fake.FirstName()
		tUser := user.FromPrimitives("", expectedName, "", "", false, false)

		if tUser.GetName() != expectedName {
			t.Fatalf("Error: %s was expected, %s was got", expectedName, tUser.GetName())
		}
	})

	t.Run("Should return user lastname", func(t *testing.T) {
		expectedLastname := fake.LastName()
		tUser := user.FromPrimitives("", "", expectedLastname, "", false, false)

		if tUser.GetLastname() != expectedLastname {
			t.Fatalf("Error: %s was expected, %s was got", expectedLastname, tUser.GetLastname())
		}
	})

	t.Run("Should return user email", func(t *testing.T) {
		expectedEmail := fake.EmailAddress()
		tUser := user.FromPrimitives("", "", "", expectedEmail, false, false)

		if tUser.GetEmail() != expectedEmail {
			t.Fatalf("Error: %s was expected, %s was got", expectedEmail, tUser.GetEmail())
		}
	})

	t.Run("Should return if user is an admin", func(t *testing.T) {
		tUser := user.FromPrimitives("", "", "", "", false, false)
		if tUser.IsAdmin() != false {
			t.Fatalf("Error: %v was expected, %v was got", false, tUser.IsAdmin())
		}
	})

	t.Run("Should return if user has email verified", func(t *testing.T) {
		tUser := user.FromPrimitives("", "", "", "", false, false)
		if tUser.HasEmailVerified() != false {
			t.Fatalf("Error: %v was expected, %v was got", false, tUser.HasEmailVerified())
		}
	})
}
