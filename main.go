package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/mattn/go-sqlite3"
	"github.com/vigneshpillutla/rssagg/api"
)


func main() {
	godotenv.Load()
	port := os.Getenv("PORT");

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	srv := &http.Server{
		Handler: api.InitRoutes(),
		Addr: ":" + port,
	}

	fmt.Println("Server is running on port: " + port)

	serverError := srv.ListenAndServe()

	if serverError!= nil {
		log.Fatal(serverError)
	}
}	