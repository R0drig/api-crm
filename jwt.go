package main
import (
	"fmt"
	"net/http"
	"time"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"github.com/dgrijalva/jwt-go"
)

var (
	db           *gorm.DB
	secretKey    = []byte("your-secret-key") // Replace with your secret key
	tokenExpires = 24 * time.Hour             // Token expiration time
)

func generateToken(user User) string {
	claims := jwt.MapClaims{
		"user_id": user.ID,
		"exp":     time.Now().Add(tokenExpires).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		fmt.Println("Error generating token:", err)
		return ""
	}

	return tokenString
}

func authenticateUser(c *gin.Context) {
	tokenString := c.GetHeader("Authorization")
	if tokenString == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Authentication required"})
		c.Abort()
		return
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if err != nil || !token.Valid {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		c.Abort()
		return
	}


	claims, _ := token.Claims.(jwt.MapClaims)
	userID := uint(claims["user_id"].(float64))

	var user User
	if err := db.First(&user, userID).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
		c.Abort()
		return
	}
	
	c.Set("user", user)
	c.Set("userID", userID)
}
