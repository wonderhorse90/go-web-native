package categorymodel

import (
	"fmt"
	"go-web-native/config"
	"go-web-native/entities"
	"log"
)

// GetAll fetches all categories
func GetAll() []entities.Category {
	var categories []entities.Category
	result := config.DB.Find(&categories)
	if result.Error != nil {
		log.Println("Error fetching categories:", result.Error)
	}
	return categories
}

// Create a new category
func Create(category entities.Category) bool {
	result := config.DB.Create(&category)
	if result.Error != nil {
		log.Println("Error creating category:", result.Error)
		return false
	}
	return result.RowsAffected > 0
}

// Detail fetches a single category by ID
func Detail(id int) *entities.Category {
	var category entities.Category
	result := config.DB.First(&category, id)
	if result.Error != nil {
		log.Println("Error fetching category details:", result.Error)
		return nil
	}
	return &category
}

// Update an existing category
func Update(id int, category entities.Category) bool {
	// Fetch the existing category
	var existingCategory entities.Category
	result := config.DB.First(&existingCategory, id)
	if result.Error != nil {
		log.Println("Error finding category to update:", result.Error)
		return false
	}

	// Update the fields
	existingCategory.Name = category.Name
	existingCategory.UpdatedAt = category.UpdatedAt

	// Save the changes
	saveResult := config.DB.Save(&existingCategory)
	if saveResult.Error != nil {
		log.Println("Error updating category:", saveResult.Error)
		return false
	}
	return saveResult.RowsAffected > 0
}

// Delete a category by ID
func Delete(id int) error {
	result := config.DB.Delete(&entities.Category{}, id)
	if result.Error != nil {
		return result.Error // Return the error if deletion fails
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("no category found with ID %d", id) // Handle cases where no rows are affected
	}

	return nil // Return nil if the deletion is successful
}
