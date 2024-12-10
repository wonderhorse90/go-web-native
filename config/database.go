package config

import (
	"log"

	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() error {
	// Open a new database connection
	dsn := "root@tcp(127.0.0.1:3306)/go_products_master?parseTime=true"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
		return err
	}

	DB = db
	log.Println("Database connected successfully")
	return nil
}

// CloseDB closes the database connection
func CloseDB() {
	if DB != nil {
		// Retrieve the underlying *sql.DB
		sqlDB, err := DB.DB()
		if err != nil {
			log.Println("Error retrieving underlying database object:", err)
			return
		}

		// Close the database connection
		err = sqlDB.Close()
		if err != nil {
			log.Println("Error closing database connection:", err)
		} else {
			log.Println("Database connection closed")
		}
	}
}
