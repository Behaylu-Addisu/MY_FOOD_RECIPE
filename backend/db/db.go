// backend/db/db.go

package db

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" // PostgreSQL driver
	"log"
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
