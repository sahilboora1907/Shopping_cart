package controllers

// Import necessary libraries
import (
	models "cart/Models"
	http "net/http"
	time "time"

	gin "github.com/gin-gonic/gin"
	jwt "github.com/golang-jwt/jwt/v5"
	bcrypt "golang.org/x/crypto/bcrypt"
)

// Function to create a user
func Signup(c *gin.Context) {
	var input models.UsersInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input!"})
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
	password, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to encrypt password!"})
		return
	}

	// Create user
	user := models.Users{Username: input.Username, Password: string(password)}
	models.DB.Create(&user)

	if user.ID == 0 {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user!"})
		return

	}

	//Respond with success message
	c.JSON(http.StatusOK, gin.H{"Message": "User created!"})
}

// Function to login a user
func Login(c *gin.Context) {
	var input models.UsersInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input!"})
		return
	}

	// Check if user exists
	var user models.Users
	models.DB.Where("username = ?", input.Username).First(&user)

	if user.Username == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid username!"})
		return
	}

	// Compare password
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid password!"})
		return
	}

	// Create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": user.Username,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	})

	// Sign token
	tokenString, err := token.SignedString([]byte("yoo-bitchhh!@1907"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token!"})
		return
	}
	// Respond with cookie
	c.SetSameSite(http.SameSiteNoneMode)
	c.SetCookie("token", tokenString, 3600*24, "", "", false, true)
	c.JSON(http.StatusOK, gin.H{})
}

func GetUsers(c *gin.Context) {
	user, _ := c.Get("user")
	c.JSON(http.StatusOK, gin.H{"data": user})
}
