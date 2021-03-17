package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	postgresdatabase "github.com/alejogs4/hn-website/src/shared/infraestructure/postgresDatabase"
	"github.com/alejogs4/hn-website/src/shared/infraestructure/server"
	"github.com/alejogs4/hn-website/src/shared/infraestructure/token"
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

	router := server.InitializeHTTPRouter(database)

	port := fmt.Sprintf(":%s", os.Getenv("PORT"))
	log.Println(fmt.Sprintf("Initializing server in port %s", port))
	log.Fatal(http.ListenAndServe(port, router))
}

func loadEnviromentVariables() {
	currentEnviroment := os.Getenv("ENV")

	if currentEnviroment == "dev" {
		godotenv.Load(".env.dev")
		return
	}

	godotenv.Load(".env")
}
