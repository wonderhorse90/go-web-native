package config

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func ConnectDB() error {
	// Open a new database connection
	db, err := sql.Open("mysql", "root@/go_products_master?parseTime=true")
	if err != nil {
		return err // return error if connection fails
	}

	// Test the connection
	if err := db.Ping(); err != nil {
		return err // return error if ping fails
	}

	// Store the DB connection in a global variable
	DB = db
	log.Println("Database connected")
	return nil
}

// CloseDB closes the database connection
func CloseDB() {
	if DB != nil {
		err := DB.Close()
		if err != nil {
			log.Println("Error closing database connection:", err)
		} else {
			log.Println("Database connection closed")
		}
	}
}
