package models

type Ingredient struct {
	ID        uint   `json:"id"`
	Name      string `json:"name"`
	Quantity  string `json:"quantity"`
	RecipeID  uint   `json:"recipe_id"`
}
