package api

import (
	"encoding/json"
	"log"
	"net/http"
)

func respondWithError(w http.ResponseWriter, code int, message string) {
	if code >= 500 {
		log.Println(message);
	}
	
	respondWithJSON(w, code, map[string]string{"error": message});
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	data, err := json.Marshal(payload);

	if err!= nil {
		w.WriteHeader(http.StatusInternalServerError);
		log.Println("Something went wrong while marshalling the data");
		return;
	}
	w.WriteHeader(code);
	w.Header().Add("Content-Type", "application/json");
	w.Write(data);
}