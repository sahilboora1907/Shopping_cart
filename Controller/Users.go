package controllers

// Import necessary libraries
import (
	models "cart/Models"
	http "net/http"
	"time"

	gin "github.com/gin-gonic/gin"
	jwt "github.com/golang-jwt/jwt/v5"
	bcrypt "golang.org/x/crypto/bcrypt"
)

// Function to create a user
func CreateUser(c *gin.Context) {
	var input models.UsersInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Check if user already exists
	var userFound models.Users
	models.DB.Where("username = ?", input.Username).First(&userFound)

	if userFound.Username != "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User already exists!"})
		return
	}
	// Encrypt password
	password, _ := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)

	// Create user
	user := models.Users{Username: input.Username, Password: string(password)}
	models.DB.Create(&user)

	c.JSON(http.StatusOK, gin.H{"data": user})
}

func Login(c *gin.Context) {
	var input models.UsersInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Check if user exists
	var user models.Users
	models.DB.Where("username = ?", input.Username).First(&user)

	if user.Username == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User not found!"})
		return
	}
	// Compare password
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid credentials!"})
		return
	}
	// Create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": user.Username,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	})
	// Sign token
	tokenString, err := token.SignedString([]byte("auth-api-jwt-secret"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": tokenString})
}
