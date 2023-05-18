package middleware

import (
	"fmt"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type AuthJSON struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

var privateKey = []byte(os.Getenv("JWT_SECRET"))

func Verify() gin.HandlerFunc {
	return func(c *gin.Context) {
		authorization := c.GetHeader("Authorization")
		tokenString := strings.Split(authorization, " ")[1]
		fmt.Println(tokenString)
		if tokenString == "null" {
			c.JSON(400, gin.H{"error": "invalid token"})
			c.Abort()
			return
		}

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("invalid token")
			}
			return privateKey, nil
		})
		if err != nil {
			c.JSON(400, gin.H{"error": "invalid token"})
			c.Abort()
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			c.JSON(400, gin.H{"error": "invalid token"})
			c.Abort()
			return
		}
		c.Set("claims", claims)
		c.Next()
		return
	}
}
