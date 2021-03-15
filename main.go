package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	postgresproductrepository "github.com/alejogs4/hn-website/src/products/infraestructure/postgresProductRepository"
	productshttpport "github.com/alejogs4/hn-website/src/products/infraestructure/productsHttpPort"
	"github.com/alejogs4/hn-website/src/shared/infraestructure/email"
	postgresdatabase "github.com/alejogs4/hn-website/src/shared/infraestructure/postgresDatabase"
	"github.com/alejogs4/hn-website/src/shared/infraestructure/token"
	userhttpport "github.com/alejogs4/hn-website/src/user/infraestructure/userHttpPort"
	userrepository "github.com/alejogs4/hn-website/src/user/infraestructure/userRepository"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	loadEnviromentVariables()
	token.LoadCertificates("./certificates/app.rsa.pub", "./certificates/app.rsa")

	database, err := postgresdatabase.LoadDatabase(
		os.Getenv("DATABASE_HOST"),
		os.Getenv("POSTGRES_DB"),
		os.Getenv("POSTGRES_USER"),
		os.Getenv("DB_PORT"),
	)
	if err != nil {
		log.Fatal(err)
	}

	httpRouter := mux.NewRouter()
	port := fmt.Sprintf(":%s", os.Getenv("PORT"))

	userhttpport.HandleUserControllers(httpRouter, userrepository.NewUserPostgresCommandsRepository(database), email.SMPTService{})
	productshttpport.HandleProductsControllers(httpRouter, postgresproductrepository.NewPostgresProductCommandsRepository(database))

	log.Println(fmt.Sprintf("Initializing server in port %s", port))
	log.Fatal(http.ListenAndServe(port, httpRouter))
}

func loadEnviromentVariables() {
	currentEnviroment := os.Getenv("ENV")

	if currentEnviroment == "dev" {
		godotenv.Load(".env.dev")
		return
	}

	godotenv.Load(".env")
}
