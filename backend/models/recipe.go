package models

type Recipe struct {
	ID             uint   `json:"id"`
	Title          string `json:"title"`
	Description    string `json:"description"`
	Ingredients    string `json:"ingredients"`
	PreparationTime int   `json:"preparation_time"`
	CreatedAt      string `json:"created_at"`
	UpdatedAt      string `json:"updated_at"`
}
