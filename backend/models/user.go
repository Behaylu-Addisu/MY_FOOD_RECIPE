package models

import (
    "github.com/google/uuid"
    "time"
)

// User represents a user in the system
type User struct {
    ID          uuid.UUID `json:"id" db:"id"`               // UUID as primary key
    Email       string    `json:"email" db:"email"`         // User's email
    Password    string    `json:"password" db:"password"`   // User's hashed password
    Name        string    `json:"name" db:"name"`           // User's name
    Username    string    `json:"username" db:"username"`   // User's username
    FirstName   string    `json:"first_name" db:"first_name"`
    LastName    string    `json:"last_name" db:"last_name"`
    ProfilePic  string    `json:"profile_picture" db:"profile_picture"`
    CreatedAt   time.Time `json:"created_at" db:"created_at"`
    UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
}
