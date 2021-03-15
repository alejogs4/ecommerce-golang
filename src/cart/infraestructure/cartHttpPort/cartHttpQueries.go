package carthttpport

import (
	"net/http"

	queryusecases "github.com/alejogs4/hn-website/src/cart/application/queryUseCases"
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
		httputils.DispatchHTTPError(rw, "Something went wrong", http.StatusInternalServerError)
		return
	}

	httputils.DispatchDefaultAPIResponse(rw, userCart, "Ok", http.StatusOK)
}
