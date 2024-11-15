package main

import (
	"fmt"
	"log"
	"os"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/Behaylu-Addisu/MY_FOOD_RECIPE/backend/handlers"
)

var db *sqlx.DB
var err error

func main() {
	// Load environment variables from .env file
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Get the database URL from the environment variable
	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		log.Fatal("DATABASE_URL environment variable not set")
	}

	// Connect to the PostgreSQL database
	db, err = sqlx.Open("postgres", dbURL)
	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}
	defer db.Close()

	// Initialize the router
	r := mux.NewRouter()

	// Define routes
	r.HandleFunc("/api/signup", handlers.SignUpHandler).Methods("POST")
	r.HandleFunc("/api/login", handlers.LoginHandler).Methods("POST")
	r.HandleFunc("/api/recipes", handlers.CreateRecipeHandler).Methods("POST")
	r.HandleFunc("/api/recipes", handlers.GetAllRecipesHandler).Methods("GET")
	r.HandleFunc("/api/recipe/{id:[0-9]+}", handlers.GetRecipeHandler).Methods("GET")
	r.HandleFunc("/api/recipe/{id:[0-9]+}", handlers.UpdateRecipeHandler).Methods("PUT")
	r.HandleFunc("/api/recipe/{id:[0-9]+}", handlers.DeleteRecipeHandler).Methods("DELETE")

	// Start the server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Default port
	}
	log.Printf("Server starting on port %s...\n", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}
