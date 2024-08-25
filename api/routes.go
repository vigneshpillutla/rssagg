package api

import (
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
)

func InitRoutes() *chi.Mux {
	createDatabaseConnection();

	router := chi.NewRouter()
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE"},
		ExposedHeaders: []string{"Link"},
		AllowedHeaders: []string{"*"},
		AllowCredentials: false,
		MaxAge: 300,
	}))

	router.Get("/healthz", readinessHandler)
	router.Get("/error", errorRoute)
	router.Post("/users", createUserHandler)

	return router;
}

func readinessHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("%v readinessHandler",r.URL)
	respondWithJSON(w, http.StatusOK, map[string]string{"status": "ok"});
}

func errorRoute(w http.ResponseWriter, r *http.Request) {
	log.Printf("%v errorRoute",r.URL)
	respondWithError(w, http.StatusBadRequest, "Something went wrong");
}