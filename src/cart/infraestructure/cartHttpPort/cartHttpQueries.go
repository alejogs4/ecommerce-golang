package carthttpport

import (
	"net/http"

	queryusecases "github.com/alejogs4/hn-website/src/cart/application/queryUseCases"
	carthttperror "github.com/alejogs4/hn-website/src/cart/infraestructure/cartHttpPort/cartHttpError"
	httputils "github.com/alejogs4/hn-website/src/shared/infraestructure/httpUtils"
	userdto "github.com/alejogs4/hn-website/src/user/application/userDTO"
)

type cartHTTPQueries struct {
	cartQueriesUseCases queryusecases.CartQueries
}

func (chq *cartHTTPQueries) GetUserCart(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-type", "application/json")

	user, _ := r.Context().Value("user").(userdto.UserLoginDTO)
	userCart, err := chq.cartQueriesUseCases.GetUserCart(user.ID)
	if err != nil {
		httpError := carthttperror.MapCartErrorToHTTPError(err)
		httputils.DispatchHTTPError(rw, httpError.Message, httpError.StatusCode)
		return
	}

	httputils.DispatchDefaultAPIResponse(rw, userCart, "Ok", http.StatusOK)
}
