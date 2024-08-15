package main

import (
	"net/http"

	"github.com/go-chi/chi"
)


func v1Handler(mainRouter *chi.Mux) {
	router := chi.NewRouter();
	router.Get("/healthz", v1ReadinessHandler)
	router.Get("/error", v1ErrorRoute)
	
	mainRouter.Mount("/v1", router);
}

func v1ReadinessHandler(w http.ResponseWriter, r *http.Request) {
	respondWithJSON(w, http.StatusOK, map[string]string{"status": "ok"});
}

func v1ErrorRoute(w http.ResponseWriter, r *http.Request) {
	responseWithError(w, http.StatusBadRequest, "Something went wrong");
}