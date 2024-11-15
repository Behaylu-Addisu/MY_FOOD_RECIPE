// db/db.go
package db

import (
    "github.com/jmoiron/sqlx"
    _ "github.com/lib/pq"
    "log"
    "github.com/Behaylu-Addisu/MY_FOOD_RECIPE/tree/main/backend/models"
)

var Db *sqlx.DB

// InitDB initializes the database connection
func InitDB(dsn string) *sqlx.DB {
    var err error
    Db, err = sqlx.Open("postgres", dsn)
    if err != nil {
        log.Fatal("Could not connect to the database:", err)
    }
    return Db
}

// CreateUser creates a new user in the database
func CreateUser(user models.User) error {
    query := `INSERT INTO users (email, password, name) VALUES ($1, $2, $3)`
    _, err := Db.Exec(query, user.Email, user.Password, user.Name)
    return err
}

// GetUserByEmail retrieves a user by email
func GetUserByEmail(email string) (models.User, error) {
    var user models.User
    query := `SELECT id, email, password, name FROM users WHERE email=$1`
    err := Db.Get(&user, query, email)
    return user, err
}

// CreateRecipe creates a new recipe in the database (if this is part of your db package)
func CreateRecipe(recipe models.Recipe) error {
    query := `INSERT INTO recipes (title, description, user_id) VALUES ($1, $2, $3)`
    _, err := Db.Exec(query, recipe.Title, recipe.Description, recipe.UserID)
    return err
}

// GetAllRecipes retrieves all recipes from the database
func GetAllRecipes() ([]models.Recipe, error) {
    var recipes []models.Recipe
    query := `SELECT id, title, description, user_id FROM recipes`
    err := Db.Select(&recipes, query)
    return recipes, err
}
