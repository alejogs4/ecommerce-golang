package userhttpport

import (
	"encoding/json"
	"net/http"

	httputils "github.com/alejogs4/hn-website/src/shared/infraestructure/httpUtils"
	"github.com/alejogs4/hn-website/src/shared/infraestructure/token"
	usercommands "github.com/alejogs4/hn-website/src/user/application/userCommands"
	usererrormapper "github.com/alejogs4/hn-website/src/user/infraestructure/userHttpPort/userErrorMapper"
)

type commandsControllers struct {
	usercommands.UseCases
}

func (cc commandsControllers) CreateUserController(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	var incommingUser struct {
		Name     string `json:"name"`
		Lastname string `json:"lastname"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	err := json.NewDecoder(r.Body).Decode(&incommingUser)
	if err != nil {
		httputils.DispatchHTTPError(w, "Something went wrong with sent data", http.StatusBadRequest)
		return
	}

	err = cc.UseCases.CreateUser(incommingUser.Name, incommingUser.Lastname, incommingUser.Email, incommingUser.Password)
	if err != nil {
		httpError := usererrormapper.MapUserErrorToHTTPError(err)
		httputils.DispatchHTTPError(w, httpError.Message, httpError.StatusCode)
		return
	}

	httputils.DispatchNewResponse(w, map[string]string{"message": "Ok"}, http.StatusCreated)
}

func (cc commandsControllers) LoginUserController(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	var loginInformation struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	err := json.NewDecoder(r.Body).Decode(&loginInformation)
	if err != nil {
		httputils.DispatchHTTPError(w, "Something went wrong getting the information", http.StatusInternalServerError)
		return
	}

	user, err := cc.UseCases.LoginUser(loginInformation.Email, loginInformation.Password)
	if err != nil {
		httpError := usererrormapper.MapUserErrorToHTTPError(err)
		httputils.DispatchHTTPError(w, httpError.Message, httpError.StatusCode)
		return
	}

	userToken, err := token.CreateToken(user)
	if err != nil {
		httpError := usererrormapper.MapUserErrorToHTTPError(err)
		httputils.DispatchHTTPError(w, httpError.Message, httpError.StatusCode)
	}

	userResponse := map[string]interface{}{"user": user, "token": userToken}
	httputils.DispatchDefaultAPIResponse(w, userResponse, "Ok", http.StatusOK)
}

func (cc commandsControllers) VerifyUserEmail(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "text/html")

	userToken := r.URL.Query().Get("token")
	userDTO, err := token.GetUserFromToken(userToken)
	w.WriteHeader(http.StatusOK)

	if err != nil {
		w.Write([]byte("<h1>Token was not found</h1>"))
		return
	}

	err = cc.UseCases.VerifyEmail(userDTO.Email)
	if err != nil {
		w.Write([]byte("<h1>Invalid token</h1>"))
		return
	}

	w.Write([]byte("<h1>Email succesfully verified</h1>"))
}
