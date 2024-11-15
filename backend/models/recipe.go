package models

import (
    "github.com/google/uuid"
    "time"
)

// Recipe represents a recipe in the system.
type Recipe struct {
    ID          uuid.UUID `json:"id" db:"id"` // UUID as primary key
    UserID      uuid.UUID `json:"user_id" db:"user_id"` // Foreign key to users table
    Title       string    `json:"title" db:"title"`
    Description string    `json:"description" db:"description"`
    CategoryID  uuid.UUID `json:"category_id" db:"category_id"` // Foreign key to categories table
    PrepTime    int       `json:"prep_time" db:"prep_time"`
    FeaturedImage string  `json:"featured_image" db:"featured_image"`
    CreatedAt   time.Time `json:"created_at" db:"created_at"`
    UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
}
