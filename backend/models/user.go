package models

import (
    "github.com/google/uuid"
    "time"
)

// User represents a user in the system.
type User struct {
    ID            uuid.UUID `json:"id" db:"id"` // UUID as primary key
    Email         string    `json:"email" db:"email"`
    PasswordHash  string    `json:"password_hash" db:"password_hash"`
    Username      string    `json:"username" db:"username"`
    FirstName     string    `json:"first_name" db:"first_name"`
    LastName      string    `json:"last_name" db:"last_name"`
    ProfilePicture string   `json:"profile_picture" db:"profile_picture"`
    CreatedAt     time.Time `json:"created_at" db:"created_at"`
    UpdatedAt     time.Time `json:"updated_at" db:"updated_at"`
}
