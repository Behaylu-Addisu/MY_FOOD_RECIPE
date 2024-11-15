package models

import (
    "github.com/google/uuid" // Correct import for the uuid package
    "time"
)

// Ingredient represents an ingredient for a recipe.
type Ingredient struct {
    ID             uuid.UUID `json:"id" db:"id"`               // UUID as primary key
    RecipeID       uuid.UUID `json:"recipe_id" db:"recipe_id"` // Foreign key to recipes table
    IngredientName string    `json:"ingredient_name" db:"ingredient_name"`
    Quantity       string    `json:"quantity" db:"quantity"`
    Unit           string    `json:"unit" db:"unit"`
    CreatedAt      time.Time `json:"created_at" db:"created_at"`
    UpdatedAt      time.Time `json:"updated_at" db:"updated_at"`
}
