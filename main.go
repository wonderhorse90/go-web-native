package main

import (
	"go-web-native/config"
	"go-web-native/controllers/categorycontroller"
	"go-web-native/controllers/homecontroller"
	"go-web-native/controllers/productcontroller"
	"go-web-native/controllers/usercontroller"
	"go-web-native/entities"
	"log"
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func main() {
	// Database connection
	config.ConnectDB()

	if err := config.ConnectDB(); err != nil {
		log.Fatal("Error connecting to database:", err)
	}
	defer config.CloseDB()

	err := config.DB.AutoMigrate(&entities.User{})
	if err != nil {
		panic("Failed to migrate database: " + err.Error())
	}

	// Initialize Gin router
	r := gin.Default()

	r.LoadHTMLGlob("views/*/*")
	templatesPath := "views/*/*"
	matchedFiles, err := filepath.Glob(templatesPath)
	if err != nil {
		log.Fatal("Error matching templates:", err)
	}
	log.Println("Matched templates:", matchedFiles)

	// Routes
	r.GET("/", func(c *gin.Context) {
		// Redirect to login page (or register page as per your flow)
		c.Redirect(http.StatusFound, "/register")
	})
	// 1. User Routes
	r.GET("/register", usercontroller.RegisterPage) // Register Page
	r.POST("/register", usercontroller.Register)    // Handle Registration
	r.GET("/login", usercontroller.LoginPage)       // Login Page
	r.POST("/login", usercontroller.Login)

	// 2. Homepage
	r.GET("/home", homecontroller.Welcome)

	// 3. Categories
	categoryGroup := r.Group("/categories")
	{
		categoryGroup.GET("/", categorycontroller.Index)
		categoryGroup.POST("/add", categorycontroller.Add)
		categoryGroup.PUT("/edit", categorycontroller.Edit)
		categoryGroup.DELETE("/delete", categorycontroller.Delete)
	}

	// 4. Products
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
