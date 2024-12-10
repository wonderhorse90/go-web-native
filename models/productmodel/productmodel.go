package productmodel

import (
	"fmt"
	"go-web-native/config"
	"go-web-native/entities"
	"log"
)

// GetAll fetches all products with their associated categories
func GetAll() []entities.Product {
	var products []entities.Product

	// Use Preload to fetch related category data
	result := config.DB.Preload("Category").Find(&products)
	if result.Error != nil {
		log.Println("Error fetching products:", result.Error)
	}

	return products
}

// Create inserts a new product into the database
func Create(product entities.Product) bool {
	result := config.DB.Create(&product)
	if result.Error != nil {
		log.Println("Error creating product:", result.Error)
		return false
	}
	return result.RowsAffected > 0
}

// Detail fetches a single product by its ID
func Detail(id int) *entities.Product {
	var product entities.Product

	// Use Preload to fetch related category data
	result := config.DB.Preload("Category").First(&product, id)
	if result.Error != nil {
		log.Println("Error fetching product details:", result.Error)
		return nil
	}

	return &product
}

// Update modifies an existing product
func Update(id int, product entities.Product) bool {
	// Fetch the existing product
	var existingProduct entities.Product
	if err := config.DB.First(&existingProduct, id).Error; err != nil {
		log.Println("Error finding product to update:", err)
		return false
	}

	// Update the fields
	existingProduct.Name = product.Name
	existingProduct.CategoryID = product.CategoryID
	existingProduct.Stock = product.Stock
	existingProduct.Description = product.Description
	existingProduct.UpdatedAt = product.UpdatedAt

	// Save changes
	saveResult := config.DB.Save(&existingProduct)
	if saveResult.Error != nil {
		log.Println("Error updating product:", saveResult.Error)
		return false
	}
	return saveResult.RowsAffected > 0
}

// Delete removes a product by its ID
func Delete(id int) error {
	result := config.DB.Delete(&entities.Product{}, id)
	if result.Error != nil {
		return result.Error // Return the error if deletion fails
	}
	if result.RowsAffected == 0 {
		return fmt.Errorf("no category found with ID %d", id) // Handle cases where no rows are affected
	}
	return nil
}
