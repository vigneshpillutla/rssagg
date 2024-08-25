package api

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/vigneshpillutla/rssagg/internal/database"
	"github.com/vigneshpillutla/rssagg/models"
)

func dbUserToAPIUser(dbUser database.User) models.User {
	return models.User{
		ID: dbUser.ID,
		CreatedAt: dbUser.CreatedAt,
		UpdatedAt: dbUser.UpdatedAt,
		Name: dbUser.Name,
	}
}

func createUserHandler(w http.ResponseWriter, r *http.Request){
	log.Printf("%v createUserHandler", r.URL);
	type parameters struct {
		Name string `json:"name"`
	}

	params := parameters{}
	err := json.NewDecoder(r.Body).Decode(&params);

	if err != nil {
		log.Printf("Error decoding the request payload: %v", err);
		respondWithError(w, http.StatusBadRequest, "Invalid request payload");
		return;
	}

	user, err := APIConfigInstance.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID: uuid.New().String(),
		Name: params.Name,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	});

	if err != nil {
		log.Printf("Error creating user: %v", err);
		respondWithError(w, http.StatusInternalServerError, "Could not create user");
		return;
	}

	respondWithJSON(w, http.StatusCreated, dbUserToAPIUser(user));
}