package handlers

import (
	"encoding/json"
	"net/http"
	"github.com/Behaylu-Addisu/MY_FOOD_RECIPE/tree/main/backend/models"
	// "github.com/Behaylu-Addisu/MY_FOOD_RECIPE/tree/main/backend/db"
)

func CreateRecipeHandler(w http.ResponseWriter, r *http.Request) {
	var recipe models.Recipe
	err := json.NewDecoder(r.Body).Decode(&recipe)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	// Save recipe to the database
	err = db.CreateRecipe(recipe)
	if err != nil {
		http.Error(w, "Failed to create recipe", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "Recipe created successfully"})
}

func GetAllRecipesHandler(w http.ResponseWriter, r *http.Request) {
	recipes, err := db.GetAllRecipes()
	if err != nil {
		http.Error(w, "Failed to fetch recipes", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(recipes)
}

func GetRecipeHandler(w http.ResponseWriter, r *http.Request) {
	// Extract recipe ID from URL
	// Fetch recipe details from the database

	// Respond with recipe details
}

func UpdateRecipeHandler(w http.ResponseWriter, r *http.Request) {
	// Update recipe logic
}

func DeleteRecipeHandler(w http.ResponseWriter, r *http.Request) {
	// Delete recipe logic
}
