package usercontroller

import (
	"go-web-native/entities"
	"go-web-native/models/usermodel"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func RegisterPage(c *gin.Context) {
	log.Println("Serving register page")
	c.HTML(http.StatusOK, "register.html", nil) // Render the registration page
}

// Register handler
func Register(c *gin.Context) {
	var user entities.User
	if err := c.ShouldBind(&user); err != nil {
		log.Println("Error binding form data:", err)
		c.JSON(400, gin.H{"error": "Missing form data"})
		return
	}

	log.Printf("Form Data: Username: %s, Email: %s, Password: %s", user.Username, user.Email, user.Password)

	// Validate that the username, email, and password are not empty
	if user.Username == "" || user.Email == "" || user.Password == "" {
		c.HTML(http.StatusBadRequest, "register.html", gin.H{
			"error": "Username, email, and password cannot be empty.",
		})
		return
	}

	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error hashing password"})
		return
	}
	user.Password = string(hashedPassword)

	newUser := entities.User{
		Username: user.Username,
		Email:    user.Email,
		Password: string(hashedPassword),
	}

	// Save user to DB
	if err := usermodel.Create(newUser); err != nil {
		c.HTML(http.StatusInternalServerError, "register.html", gin.H{
			"error": "Error creating user, please try again.",
		})
		return
	}

	c.Redirect(http.StatusFound, "/login")
}

func LoginPage(c *gin.Context) {
	log.Println("Serving login page")
	c.HTML(http.StatusOK, "login.html", nil) // Render the login page
}

// Login handler
func Login(c *gin.Context) {
	var user entities.User
	if err := c.ShouldBind(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Fetch user from DB
	storedUser, err := usermodel.GetByUsername(user.Username)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	// Compare passwords
	if err := bcrypt.CompareHashAndPassword([]byte(storedUser.Password), []byte(user.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	// Create session or JWT token

	c.Redirect(http.StatusFound, "/home")
}
