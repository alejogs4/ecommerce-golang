package productshttpport

import (
	productcommands "github.com/alejogs4/hn-website/src/products/application/productCommands"
	"github.com/alejogs4/hn-website/src/products/domain/product"
	"github.com/alejogs4/hn-website/src/shared/infraestructure/authorization"
	"github.com/alejogs4/hn-website/src/shared/infraestructure/middleware"
	"github.com/gorilla/mux"
)

// HandleProductsControllers build the routes for products
func HandleProductsControllers(router *mux.Router, productsCommandsRepository product.CommandsRepository) {
	productsCommandsUseCases := productcommands.NewProductCommandsUseCases(productsCommandsRepository)
	productCommandsController := productCommandsControllers{ProductCommandUseCases: productsCommandsUseCases}

	router.HandleFunc("/api/v1/product", middleware.Chain(
		productCommandsController.CreateProductController,
		authorization.EnsureEmailVerified(),
		authorization.AdminOperation(),
		authorization.ValidateToken(),
	))
}
