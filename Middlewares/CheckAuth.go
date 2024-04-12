package middleware

import (
	models "cart/Models"
	fmt "fmt"
	http "net/http"
	time "time"

	gin "github.com/gin-gonic/gin"
	jwt "github.com/golang-jwt/jwt/v5"
)

// CheckAuth is a middleware that checks if the request contains a valid Authorization header
func CheckAuth(c *gin.Context) {

	//Get the Cookie
	cookie, err := c.Cookie("token")
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "No token found"})
		c.AbortWithStatus(http.StatusUnauthorized)
		return

	}

	// Check if the token is valid
	token, err := jwt.Parse(cookie, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])

		}
		return []byte("yoo-bitchhh!@1907"), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {

		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token expired"})
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		// Find the user in the database
		var user models.Users
		models.DB.Where("username = ?", claims["username"]).First(&user)

		if user.Username == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		// Set the user in the context
		c.Set("user", user)
		c.Next()
	}
}
