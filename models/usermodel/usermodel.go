package usermodel

import (
	"errors"
	"go-web-native/config"
	"go-web-native/entities"

	"log"
)

// Create a new user in the database
func Create(user entities.User) error {
	// Use the db connection from config
	result := config.DB.Create(&user)
	return result.Error
}

// Get a user by username
func GetByUsername(username string) (entities.User, error) {
	var user entities.User
	result := config.DB.Where("username = ?", username).First(&user)
	if result.Error != nil {
		return user, errors.New("user not found")
	}
	return user, nil
}

// Get a user by email
func GetByEmail(email string) *entities.User {
	var user entities.User
	result := config.DB.Where("email = ?", email).First(&user)
	if result.Error != nil {
		log.Println("Error fetching user by email:", result.Error)
		return nil
	}
	return &user
}
