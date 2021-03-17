package userhttpport_test

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/alejogs4/hn-website/src/infraestructure/mocks/mock_mailservice"
	"github.com/alejogs4/hn-website/src/infraestructure/mocks/mock_user"
	"github.com/alejogs4/hn-website/src/shared/infraestructure/server"
	"github.com/alejogs4/hn-website/src/shared/infraestructure/token"
	"github.com/golang/mock/gomock"
	"github.com/icrowley/fake"
)

func TestSignupEndpoint(t *testing.T) {
	t.Run("Should return 201 if user was succesfully created", func(t *testing.T) {
		t.Parallel()
		err := token.LoadCertificates("../../../../certificates/app.rsa.pub", "../../../../certificates/app.rsa")

		if err != nil {
			t.Fatal(err)
		}
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
		emailService.EXPECT().Send(gomock.Any()).Return(nil)

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
}
