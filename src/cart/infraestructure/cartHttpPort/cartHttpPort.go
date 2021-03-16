package carthttpport

import (
	"net/http"

	commandusecases "github.com/alejogs4/hn-website/src/cart/application/commandUseCases"
	queryusecases "github.com/alejogs4/hn-website/src/cart/application/queryUseCases"
	"github.com/alejogs4/hn-website/src/cart/domain/cart"
	cartevents "github.com/alejogs4/hn-website/src/cart/domain/cart/cartEvents"
	cartemail "github.com/alejogs4/hn-website/src/cart/infraestructure/cartEmail"
	"github.com/alejogs4/hn-website/src/products/domain/product"
	productevents "github.com/alejogs4/hn-website/src/products/infraestructure/productEvents"
	"github.com/alejogs4/hn-website/src/shared/infraestructure/authorization"
	mailservice "github.com/alejogs4/hn-website/src/shared/infraestructure/email/mailService"
	httputils "github.com/alejogs4/hn-website/src/shared/infraestructure/httpUtils"
	"github.com/alejogs4/hn-website/src/shared/infraestructure/middleware"
	"github.com/alejogs4/hn-website/src/user/domain/user"
	"github.com/gorilla/mux"
)

func HandleCartRoutes(
	router *mux.Router,
	queries cart.QueriesRepository,
	commands cart.CommandsRepository,
	productQueries product.QueriesRepository,
	productsCommands product.CommandsRepository,
	emailService mailservice.Service,
	userQueries user.Queries,
) {

	cartQueries := cartHTTPQueries{cartQueriesUseCases: queryusecases.NewCartQueries(queries)}

	// create car use cases and register its events
	cartUseCases := commandusecases.NewCartCommand(commands, queries, productQueries)
	cartUseCases.RegisterEventHandler(cartevents.Bougth, cartemail.AfterBoughtCart{EmailService: emailService, UserQueries: userQueries})
	cartUseCases.RegisterEventHandler(cartevents.Bougth, productevents.AfterCarWasOrdered{ProductCommands: productsCommands})

	// add use cases
	cartCommands := cartHTTPCommands{cartCommands: cartUseCases}

	router.HandleFunc("/api/v1/user/cart", middleware.Chain(
		cartQueries.GetUserCart,
		httputils.Method(http.MethodGet),
		authorization.ValidateToken(),
		authorization.EnsureEmailVerified(),
	))

	router.HandleFunc("/api/v1/user/cart/buy", middleware.Chain(
		cartCommands.BuyCartController,
		httputils.Method(http.MethodPost),
		authorization.ValidateToken(),
		authorization.EnsureEmailVerified(),
	))

	router.Handle("/api/v1/cart/product", middleware.Chain(
		cartCommands.AddCartItem,
		httputils.Method(http.MethodPost),
		authorization.ValidateToken(),
		authorization.EnsureEmailVerified(),
	))

	// router.

	router.Handle("/api/v1/cart/{cartID}/item/{cartItem}", middleware.Chain(
		cartCommands.RemoverCartItemController,
		httputils.Method(http.MethodDelete),
		authorization.ValidateToken(),
		authorization.EnsureEmailVerified(),
	))
}
