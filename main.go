package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	port := os.Getenv("PORT");

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	router := chi.NewRouter();
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE"},
		ExposedHeaders: []string{"Link"},
		AllowedHeaders: []string{"*"},
		AllowCredentials: false,
		MaxAge: 300,
	}))

	v1Handler(router);

	srv := &http.Server{
		Handler: router,
		Addr: ":" + port,
	}

	fmt.Println("Server is running on port: " + port)

	err := srv.ListenAndServe()

	if err!= nil {
		log.Fatal(err)
	}
}	