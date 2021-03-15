package productshttpport

import (
	"net/http"

	"github.com/alejogs4/hn-website/src/products/application/productquery"
	productshttperror "github.com/alejogs4/hn-website/src/products/infraestructure/productsHttpPort/productsHttpError"
	httputils "github.com/alejogs4/hn-website/src/shared/infraestructure/httpUtils"
	"github.com/gorilla/mux"
)

type productQueriesControllers struct {
	productsUseCases productquery.ProductQueriesUseCases
}

func (pqc *productQueriesControllers) GetAllController(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-type", "application/json")

	products, err := pqc.productsUseCases.GetAllProducts()
	if err != nil {
		httpError := productshttperror.MapProductErrorToHTTPError(err)
		httputils.DispatchHTTPError(rw, httpError.Message, httpError.StatusCode)
		return
	}

	httputils.DispatchDefaultAPIResponse(rw, products, "Ok", http.StatusOK)
}

func (pqc *productQueriesControllers) GetByIDController(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-type", "application/json")

	productID := mux.Vars(r)["id"]

	product, err := pqc.productsUseCases.GetProductByID(productID)
	if err != nil {
		httpError := productshttperror.MapProductErrorToHTTPError(err)
		httputils.DispatchHTTPError(rw, httpError.Message, httpError.StatusCode)
		return
	}

	httputils.DispatchDefaultAPIResponse(rw, product, "Ok", http.StatusOK)
}
