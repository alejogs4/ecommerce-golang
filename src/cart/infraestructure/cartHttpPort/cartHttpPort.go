package carthttpport

import (
	"net/http"

	commandusecases "github.com/alejogs4/hn-website/src/cart/application/commandUseCases"
	queryusecases "github.com/alejogs4/hn-website/src/cart/application/queryUseCases"
	"github.com/alejogs4/hn-website/src/cart/domain/cart"
	"github.com/alejogs4/hn-website/src/products/domain/product"
	"github.com/alejogs4/hn-website/src/shared/infraestructure/authorization"
	httputils "github.com/alejogs4/hn-website/src/shared/infraestructure/httpUtils"
	"github.com/alejogs4/hn-website/src/shared/infraestructure/middleware"
	"github.com/gorilla/mux"
)

func HandleCartRoutes(
	router *mux.Router,
	queries cart.QueriesRepository,
	commands cart.CommandsRepository,
	productQueries product.QueriesRepository,
) {
	cartQueries := cartHTTPQueries{cartQueriesUseCases: queryusecases.NewCartQueries(queries)}
	cartCommands := cartHTTPCommands{cartCommands: commandusecases.NewCartCommand(commands, queries, productQueries)}

	router.HandleFunc("/api/v1/user/cart", middleware.Chain(
		cartQueries.GetUserCart,
		httputils.Method(http.MethodGet),
		authorization.ValidateToken(),
		authorization.EnsureEmailVerified(),
	))

	router.Handle("/api/v1/cart/product", middleware.Chain(
		cartCommands.AddCartItem,
		httputils.Method(http.MethodPost),
		authorization.ValidateToken(),
		authorization.EnsureEmailVerified(),
	))
}
