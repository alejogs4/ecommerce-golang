package productshttpport

import (
	"encoding/json"
	"net/http"

	productcommands "github.com/alejogs4/hn-website/src/products/application/productCommands"
	productshttperror "github.com/alejogs4/hn-website/src/products/infraestructure/productsHttpPort/productsHttpError"
	httputils "github.com/alejogs4/hn-website/src/shared/infraestructure/httpUtils"
)

// ProductCommandsControllers .
type productCommandsControllers struct {
	productcommands.ProductCommandUseCases
}

// CreateProductController receives product and try to create one if user is an admin with a verified email
func (pcc *productCommandsControllers) CreateProductController(w http.ResponseWriter, r *http.Request) {
	r.Header.Set("Content-type", "application/json")
	var productInformation struct {
		Name        string  `json:"name"`
		Description string  `json:"description"`
		Picture     string  `json:"picture"`
		Quantity    int     `json:"quantity"`
		Price       float64 `json:"price"`
	}

	if err := json.NewDecoder(r.Body).Decode(&productInformation); err != nil {
		httpError := productshttperror.MapProductErrorToHTTPError(err)
		httputils.DispatchHTTPError(w, httpError.Message, httpError.StatusCode)
		return
	}

	productID, err := pcc.CreateNewProduct(
		productInformation.Name,
		productInformation.Description,
		productInformation.Picture,
		productInformation.Quantity,
		productInformation.Price,
	)
	if err != nil {
		httpError := productshttperror.MapProductErrorToHTTPError(err)
		httputils.DispatchHTTPError(w, httpError.Message, httpError.StatusCode)
		return
	}

	httputils.DispatchDefaultAPIResponse(w, productID, "Product created", http.StatusCreated)
}
