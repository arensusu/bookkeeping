package controller

import (
	"bookkeeping-backend/helper"
	"bookkeeping-backend/model"
	"fmt"

	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type AuthJSON struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

var privateKey = []byte("secret")

func GetUser(c *gin.Context) {
	username := c.MustGet("claims").(jwt.MapClaims)["username"].(string)
	if _, err := model.GetUserByName(username); err != nil {
		c.JSON(404, gin.H{"error": "user not found"})
		return
	}
	c.JSON(200, gin.H{"username": username})
	return
}

func Register(c *gin.Context) {
	var input AuthJSON
	if err := c.BindJSON(&input); err != nil {
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(400, gin.H{"error": "unable to create user"})
		return
	}

	user := model.User{
		Username: strings.TrimSpace(input.Username),
		Password: string(hashedPassword),
	}

	if err := model.CreateUser(&user); err != nil {
		c.JSON(400, gin.H{"error": "unable to create user"})
		return
	}
	jwtToken := helper.GenerateJWT(user.Username)
	c.JSON(201, gin.H{"jwt": jwtToken})
	return
}

func Login(c *gin.Context) {
	var input AuthJSON
	if err := c.BindJSON(&input); err != nil {
		return
	}

	user, err := model.GetUserByName(input.Username)
	if err != nil {
		c.JSON(404, gin.H{"error": "user not found"})
		return
	}
	fmt.Println(user)

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password))
	if err != nil {
		c.JSON(400, gin.H{"error": "invalid credentials"})
		return
	}

	jwtToken := helper.GenerateJWT(user.Username)
	c.JSON(200, gin.H{"jwt": jwtToken})
	return
}

func DeleteUser(c *gin.Context) {
	claims := c.MustGet("claims").(jwt.MapClaims)
	username := claims["username"].(string)

	user, err := model.GetUserByName(username)
	if err != nil {
		c.JSON(404, gin.H{"error": "user not found"})
		return
	}

	if err := model.DeleteUser(user.ID); err != nil {
		c.JSON(400, gin.H{"error": "unable to delete user"})
		return
	}
	c.JSON(200, gin.H{"message": "user deleted"})
	return
}
