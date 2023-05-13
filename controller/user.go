package controller

import (
	"firstapp/helper"
	"firstapp/model"
	"fmt"

	"strings"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type AuthJSON struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

var privateKey = []byte("secret")

func GetUser(c *gin.Context) {
	id := c.Param("id")
	user, err := model.GetUserById(id)
	if err != nil {
		c.JSON(404, gin.H{"error": "user not found"})
		return
	}
	fmt.Println(user)
	c.JSON(200, user)
}

func Register(c *gin.Context) {
	var input AuthJSON
	if err := c.BindJSON(&input); err != nil {
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(400, gin.H{"error": "unable to create user"})
	}

	user := model.User{
		Username: strings.TrimSpace(input.Username),
		Password: string(hashedPassword),
	}

	createdUser, err := user.Save()
	if err != nil {
		c.JSON(400, gin.H{"error": "unable to create user"})
	}
	jwtToken := helper.GenerateJWT(createdUser.Username)
	c.JSON(201, gin.H{"jwt": jwtToken})
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
}
