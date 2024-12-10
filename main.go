package main

import (
	"go-web-native/config"
	"go-web-native/controllers/categorycontroller"
	"go-web-native/controllers/homecontroller"
	"go-web-native/controllers/productcontroller"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	// Database connection
	config.ConnectDB()

	if err := config.ConnectDB(); err != nil {
		log.Fatal("Error connecting to database:", err)
	}
	defer config.CloseDB()

	// Initialize Gin router
	r := gin.Default()

	r.LoadHTMLGlob("views/*")

	// Routes
	// 1. Homepage
	r.GET("/", homecontroller.Welcome)

	// 2. Categories
	categoryGroup := r.Group("/categories")
	{
		categoryGroup.GET("/", categorycontroller.Index)
		categoryGroup.POST("/add", categorycontroller.Add)
		categoryGroup.PUT("/edit", categorycontroller.Edit)
		categoryGroup.DELETE("/delete", categorycontroller.Delete)
	}

	// 3. Products
	productGroup := r.Group("/products")
	{
		productGroup.GET("/", productcontroller.Index)
		productGroup.POST("/add", productcontroller.Add)
		productGroup.GET("/detail", productcontroller.Detail)
		productGroup.PUT("/edit", productcontroller.Edit)
		productGroup.DELETE("/delete", productcontroller.Delete)
	}

	// Run server
	log.Println("Server running on port: 8080")
	log.Fatal(r.Run(":8080"))
}
