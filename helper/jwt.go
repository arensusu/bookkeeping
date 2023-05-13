package helper

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var privateKey = []byte(os.Getenv("JWT_SECRET"))

func GenerateJWT(username string) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	})
	jwtToken, err := token.SignedString(privateKey)
	if err != nil {
		return ""
	}
	return jwtToken
}
