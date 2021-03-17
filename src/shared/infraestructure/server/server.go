package server

import (
	"database/sql"

	"github.com/alejogs4/hn-website/src/cart/domain/cart"
	carthttpport "github.com/alejogs4/hn-website/src/cart/infraestructure/cartHttpPort"
	cartpostgresrepository "github.com/alejogs4/hn-website/src/cart/infraestructure/cartPostgresRepository"
	"github.com/alejogs4/hn-website/src/products/domain/product"
	postgresproductrepository "github.com/alejogs4/hn-website/src/products/infraestructure/postgresProductRepository"
	productshttpport "github.com/alejogs4/hn-website/src/products/infraestructure/productsHttpPort"
	"github.com/alejogs4/hn-website/src/shared/infraestructure/email"
	mailservice "github.com/alejogs4/hn-website/src/shared/infraestructure/email/mailService"
	"github.com/alejogs4/hn-website/src/user/domain/user"
	userhttpport "github.com/alejogs4/hn-website/src/user/infraestructure/userHttpPort"
	userrepository "github.com/alejogs4/hn-website/src/user/infraestructure/userRepository"
	"github.com/gorilla/mux"
)

// Options for upload this application http router
type Options struct {
	Database                  *sql.DB
	UserCommandsRepository    user.CommandsRepository
	UserQueries               user.Queries
	EmailService              mailservice.Service
	ProductsQueriesRepository product.QueriesRepository
	ProductsCommands          product.CommandsRepository
	CartQueries               cart.QueriesRepository
	CartCommands              cart.CommandsRepository
}

// Option modifier functions in order set new values apart from the default ones
type Option func(*Options)

// WithDatabase functional setter function for sql database
func WithDatabase(database *sql.DB) Option {
	return func(so *Options) {
		so.Database = database
	}
}

// WithUserCommandsRepository functional setter function for user.CommandsRepository
func WithUserCommandsRepository(userCommandsRepository user.CommandsRepository) Option {
	return func(so *Options) {
		so.UserCommandsRepository = userCommandsRepository
	}
}

// WithUserQueries functional setter function for user.Queries
func WithUserQueries(userQueries user.Queries) Option {
	return func(so *Options) {
		so.UserQueries = userQueries
	}
}

// WithEmailService functional setter function for mailservice.Service
func WithEmailService(emailService mailservice.Service) Option {
	return func(so *Options) {
		so.EmailService = emailService
	}
}

// WithProductsQueriesRepository functional setter function for product.QueriesRepository
func WithProductsQueriesRepository(productsQueriesRepository product.QueriesRepository) Option {
	return func(so *Options) {
		so.ProductsQueriesRepository = productsQueriesRepository
	}
}

// WithProductsCommands functional setter function for product.CommandsRepository
func WithProductsCommands(productsCommands product.CommandsRepository) Option {
	return func(so *Options) {
		so.ProductsCommands = productsCommands
	}
}

// WithCartQueries functional setter function for cart.QueriesRepository
func WithCartQueries(cartQueries cart.QueriesRepository) Option {
	return func(so *Options) {
		so.CartQueries = cartQueries
	}
}

// WithCartCommands functional setter function for cart.CommandsRepository
func WithCartCommands(cartCommands cart.CommandsRepository) Option {
	return func(so *Options) {
		so.CartCommands = cartCommands
	}
}

// InitializeHTTPRouter .
func InitializeHTTPRouter(database *sql.DB, handlers ...Option) *mux.Router {
	httpRouter := mux.NewRouter()

	option := &Options{
		Database:                  database,
		ProductsQueriesRepository: postgresproductrepository.NewProductQueriesPostgresRepository(database),
		UserQueries:               userrepository.NewUserRepositoryPostgresQueries(database),
		UserCommandsRepository:    userrepository.NewUserPostgresCommandsRepository(database),
		EmailService:              email.SMPTService{},
		ProductsCommands:          postgresproductrepository.NewPostgresProductCommandsRepository(database),
		CartQueries:               cartpostgresrepository.NewCartQueriesPostgresRespository(database),
		CartCommands:              cartpostgresrepository.NewCartCommandsPostgresRespository(database),
	}

	for _, optionHandler := range handlers {
		optionHandler(option)
	}

	userhttpport.HandleUserControllers(
		httpRouter,
		option.UserCommandsRepository,
		option.EmailService,
	)

	productshttpport.HandleProductsControllers(
		httpRouter,
		option.ProductsCommands,
		option.ProductsQueriesRepository,
	)

	carthttpport.HandleCartRoutes(
		httpRouter,
		option.CartQueries,
		option.CartCommands,
		option.ProductsQueriesRepository,
		option.ProductsCommands,
		option.EmailService,
		option.UserQueries,
	)
	return httpRouter
}
