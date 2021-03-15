package productshttpport

import (
	"net/http"

	productcommands "github.com/alejogs4/hn-website/src/products/application/productCommands"
	"github.com/alejogs4/hn-website/src/products/application/productquery"
	"github.com/alejogs4/hn-website/src/products/domain/product"
	"github.com/alejogs4/hn-website/src/shared/infraestructure/authorization"
	httputils "github.com/alejogs4/hn-website/src/shared/infraestructure/httpUtils"
	"github.com/alejogs4/hn-website/src/shared/infraestructure/middleware"
	"github.com/gorilla/mux"
)

// HandleProductsControllers build the routes for products
func HandleProductsControllers(
	router *mux.Router,
	productsCommandsRepository product.CommandsRepository,
	productsQueriesRepository product.QueriesRepository,
) {
	productsCommandsUseCases := productcommands.NewProductCommandsUseCases(productsCommandsRepository)
	productCommandsController := productCommandsControllers{ProductCommandUseCases: productsCommandsUseCases}
	productsQueriesController := productQueriesControllers{productsUseCases: productquery.NewProductQueriesUseCases(productsQueriesRepository)}

	router.HandleFunc("/api/v1/product", middleware.Chain(
		productCommandsController.CreateProductController,
		httputils.Method(http.MethodPost),
		authorization.ValidateToken(),
		authorization.EnsureEmailVerified(),
		authorization.AdminOperation(),
	))

	router.HandleFunc("/api/v1/products", middleware.Chain(
		productsQueriesController.GetAllController,
		httputils.Method(http.MethodGet),
	))

	router.HandleFunc("/api/v1/product/:id", middleware.Chain(
		productsQueriesController.GetByIDController,
		httputils.Method(http.MethodGet),
	))
}
