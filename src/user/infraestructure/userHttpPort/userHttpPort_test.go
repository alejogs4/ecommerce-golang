package userhttpport_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/alejogs4/hn-website/src/infraestructure/mocks/mock_mailservice"
	"github.com/alejogs4/hn-website/src/infraestructure/mocks/mock_user"
	"github.com/alejogs4/hn-website/src/shared/infraestructure/server"
	"github.com/alejogs4/hn-website/src/shared/infraestructure/token"
	"github.com/alejogs4/hn-website/src/user/domain/user"
	"github.com/golang/mock/gomock"
	"github.com/icrowley/fake"
)

func TestMain(t *testing.M) {
	err := token.LoadCertificates("../../../../certificates/app.rsa.pub", "../../../../certificates/app.rsa")
	if err != nil {
		log.Fatalf("Error loading certificates %s", err)
		os.Exit(1)
		return
	}

	t.Run()
}

func TestSignupEndpoint(t *testing.T) {
	t.Run("Should return 201 if user was succesfully created", func(t *testing.T) {
		t.Parallel()

		db, _, err := sqlmock.New()
		if err != nil {
			t.Fatalf("Error: error meanwhile connectig with db stub - %v", err)
		}

		userMockController := gomock.NewController(t)
		defer userMockController.Finish()

		emailMockController := gomock.NewController(t)
		defer emailMockController.Finish()

		userCommandsRepository := mock_user.NewMockCommandsRepository(userMockController)
		userCommandsRepository.EXPECT().CreateUser(gomock.Any()).Return(nil)

		emailService := mock_mailservice.NewMockService(emailMockController)
		emailService.EXPECT().Send(gomock.Any()).Return(nil).AnyTimes()

		router := server.InitializeHTTPRouter(
			db,
			server.WithUserCommandsRepository(userCommandsRepository),
			server.WithEmailService(emailService),
		)

		mockServer := httptest.NewServer(router)
		defer mockServer.Close()
		path := "/api/v1/signup"
		petitionURL := fmt.Sprintf("%s%s", mockServer.URL, path)

		userSignupBody := []byte(fmt.Sprintf(
			`{"email": "%v", "name": "%v", "lastname": "%v", "password": "%v"}`,
			fake.EmailAddress(),
			fake.FirstName(),
			fake.LastName(),
			fake.Password(7, 16, true, true, true),
		))

		resp, err := http.Post(petitionURL, "application/json", bytes.NewBuffer(userSignupBody))
		if err != nil {
			t.Fatalf("Error: error %v was not expected", err)
		}

		if resp.StatusCode != http.StatusCreated {
			t.Fatalf("Error: Expected status code %d, Got status code %d", http.StatusCreated, resp.StatusCode)
		}
	})

	t.Run("Should return 400 if user information is not complete", func(t *testing.T) {
		t.Parallel()

		db, _, err := sqlmock.New()
		if err != nil {
			t.Fatalf("Error: error meanwhile connectig with db stub - %v", err)
		}

		userMockController := gomock.NewController(t)
		defer userMockController.Finish()

		emailMockController := gomock.NewController(t)
		defer emailMockController.Finish()

		userCommandsRepository := mock_user.NewMockCommandsRepository(userMockController)
		// The error must occurred in user domain entity rather than in repository
		userCommandsRepository.EXPECT().CreateUser(gomock.Any()).Return(nil).Times(0)

		emailService := mock_mailservice.NewMockService(emailMockController)
		emailService.EXPECT().Send(gomock.Any()).Return(nil).Times(0)

		router := server.InitializeHTTPRouter(
			db,
			server.WithUserCommandsRepository(userCommandsRepository),
			server.WithEmailService(emailService),
		)

		mockServer := httptest.NewServer(router)
		defer mockServer.Close()

		path := "/api/v1/signup"
		apiEndpoint := fmt.Sprintf("%s%s", mockServer.URL, path)

		signupBody := []byte(fmt.Sprintf(
			`{"email": "%v", "name": "", "lastname": "%v", "password": "%v"}`,
			fake.EmailAddress(), fake.LastName(), fake.Password(6, 10, true, true, true)))

		resp, err := http.Post(apiEndpoint, "application/json", bytes.NewBuffer(signupBody))
		if err != nil {
			t.Fatalf("Error: error sending signup petition - %v", err)
		}

		if resp.StatusCode != http.StatusBadRequest {
			t.Fatalf("Error: Expected status code %d, Got status code %d", http.StatusBadRequest, resp.StatusCode)
		}
	})
}

func TestLoginEndpoint(t *testing.T) {
	t.Run("Should return a 200 if user is correctly and a valid token", func(t *testing.T) {
		t.Parallel()

		db, _, err := sqlmock.New()
		if err != nil {
			t.Fatalf("Error: error meanwhile connectig with db stub - %v", err)
		}

		rightEmail := fake.EmailAddress()
		rightPassword := fake.Password(7, 10, true, true, true)
		loggedUser, _ := user.NewUser("id", fake.FirstName(), fake.LastName(), rightEmail, rightPassword, false, false)

		userCommandsController := gomock.NewController(t)
		defer userCommandsController.Finish()

		emailServiceController := gomock.NewController(t)
		defer emailServiceController.Finish()

		userCommandsMock := mock_user.NewMockCommandsRepository(userCommandsController)
		userCommandsMock.EXPECT().LoginUser(rightEmail, rightPassword).Return(loggedUser, nil)

		emailService := mock_mailservice.NewMockService(emailServiceController)
		emailService.EXPECT().Send(gomock.Any()).AnyTimes()

		router := server.InitializeHTTPRouter(
			db,
			server.WithUserCommandsRepository(userCommandsMock),
			server.WithEmailService(emailService),
		)

		mockServer := httptest.NewServer(router)
		defer mockServer.Close()

		loginEndpoint := fmt.Sprintf("%s/api/v1/login", mockServer.URL)
		loginBody := []byte(fmt.Sprintf(
			`{"email": "%v", "password": "%v"}`,
			rightEmail,
			rightPassword,
		))

		resp, err := http.Post(loginEndpoint, "application/json", bytes.NewBuffer(loginBody))
		if err != nil {
			t.Fatalf("Error: error sending login petition - %v", err)
		}

		if resp.StatusCode != http.StatusOK {
			t.Fatalf("Error: Expected status code %d, Got status code %d", http.StatusOK, resp.StatusCode)
		}

		var loginResponse struct {
			Data struct {
				Token string `json:"token"`
			} `json:"data"`
		}

		err = json.NewDecoder(resp.Body).Decode(&loginResponse)
		if err != nil {
			t.Fatalf("Error: error getting login token - %v", err)
		}

		gotUser, err := token.GetUserFromToken(loginResponse.Data.Token)
		if err != nil {
			t.Fatalf("Error: invalid generated token - %v", err)
		}

		if gotUser.Email != loggedUser.GetEmail() {
			t.Fatalf("Error: Expected user email %s, Got user email %v", loggedUser.GetEmail(), gotUser.Email)
		}
	})

	t.Run("Should return 400 if user email or password aren't right", func(t *testing.T) {
		t.Parallel()

		db, _, err := sqlmock.New()
		if err != nil {
			t.Fatalf("Error: error meanwhile connectig with db stub - %v", err)
		}

		userCommandsController := gomock.NewController(t)
		defer userCommandsController.Finish()

		emailServiceController := gomock.NewController(t)
		defer emailServiceController.Finish()

		wrongEmail := fake.EmailAddress()
		wrongPassword := fake.Password(6, 14, true, true, true)

		userCommandsMockRepository := mock_user.NewMockCommandsRepository(userCommandsController)
		userCommandsMockRepository.EXPECT().LoginUser(wrongEmail, wrongPassword).Return(user.User{}, user.ErrInvalidUser)

		emailMockService := mock_mailservice.NewMockService(emailServiceController)
		emailMockService.EXPECT().Send(gomock.Any()).Return(nil).AnyTimes()

		router := server.InitializeHTTPRouter(
			db,
			server.WithUserCommandsRepository(userCommandsMockRepository),
			server.WithEmailService(emailMockService),
		)

		mockServer := httptest.NewServer(router)
		defer mockServer.Close()

		apiEndpoint := fmt.Sprintf("%s/api/v1/login", mockServer.URL)
		loginBody := []byte(fmt.Sprintf(
			`{"email": "%v", "password": "%v"}`,
			wrongEmail,
			wrongPassword,
		))

		resp, err := http.Post(apiEndpoint, "application/json", bytes.NewBuffer(loginBody))
		if err != nil {
			t.Fatalf("Error: error sending login petition - %v", err)
		}

		if resp.StatusCode != http.StatusBadRequest {
			t.Fatalf("Error: Expected status code %d, Got status code %d", http.StatusBadRequest, resp.StatusCode)
		}
	})
}
