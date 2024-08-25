package api

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
	"github.com/vigneshpillutla/rssagg/internal/database"
)

type APIConfig struct {
	DB *database.Queries
}

var APIConfigInstance *APIConfig = &APIConfig{}


func createDatabaseConnection() {
	dbUrl := os.Getenv("DB_URL");
	dbEngine := os.Getenv("DB_ENGINE");

	if dbUrl == "" || dbEngine == "" {
		log.Fatal("$DB_URL and $DB_ENGINE must be set")
	}

	conn, err := sql.Open(dbEngine, dbUrl);

	if err != nil {
		log.Fatal("Cannot connect to database")
	}

	log.Println("Connected to the database")

	APIConfigInstance.DB = database.New(conn)
}