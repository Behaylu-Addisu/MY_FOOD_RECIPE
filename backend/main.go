package main

import (
	"fmt"
	"log"
	"os"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"net/http"
	"github.com/Behaylu-Addisu/MY_FOOD_RECIPE/backend/handlers"
)

var db *sqlx.DB
var err error

func main() {
	// Load environment variables
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Connect to PostgreSQL database
	dbURL := os.Getenv("postgres://behaylu:ngst@postgres:5432/food_recipe_db")
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
		port = "8080"
	}
	log.Fatal(http.ListenAndServe(":"+port, r))
}
