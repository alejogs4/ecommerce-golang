package carthttpport

import (
	"encoding/json"
	"net/http"

	commandusecases "github.com/alejogs4/hn-website/src/cart/application/commandUseCases"
	httputils "github.com/alejogs4/hn-website/src/shared/infraestructure/httpUtils"
	userdto "github.com/alejogs4/hn-website/src/user/application/userDTO"
	"github.com/gorilla/mux"
)

type cartHTTPCommands struct {
	cartCommands commandusecases.CartCommands
}

func (ctc *cartHTTPCommands) AddCartItem(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-type", "application/json")

	var cartItem struct {
		ProductID string `json:"product_id"`
	}

	if err := json.NewDecoder(r.Body).Decode(&cartItem); err != nil {
		httputils.DispatchHTTPError(rw, "Something went wrong", http.StatusInternalServerError)
		return
	}

	loggedUser, _ := r.Context().Value("user").(userdto.UserLoginDTO)
	cartItemID, err := ctc.cartCommands.AddCartItem(loggedUser.ID, cartItem.ProductID)

	if err != nil {
		httputils.DispatchHTTPError(rw, "Something went wrong", http.StatusInternalServerError)
		return
	}

	httputils.DispatchDefaultAPIResponse(rw, cartItemID, "item added", http.StatusCreated)
}

func (ctc *cartHTTPCommands) RemoverCartItemController(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-type", "application/json")

	cartItemID := mux.Vars(r)["cartItem"]
	cartID := mux.Vars(r)["cartID"]
	loggedUser, _ := r.Context().Value("user").(userdto.UserLoginDTO)

	err := ctc.cartCommands.RemoveCartItem(loggedUser.ID, cartItemID, cartID)
	if err != nil {
		httputils.DispatchHTTPError(rw, "Something went wrong", http.StatusInternalServerError)
		return
	}

	httputils.DispatchDefaultAPIResponse(rw, cartItemID, "item removed", http.StatusOK)
}
