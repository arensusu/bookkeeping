package controller

import (
	"firstapp/model"
	"fmt"

	"github.com/gin-gonic/gin"
)

func GetCategories(c *gin.Context) {
	categories, err := model.GetAllCategories()
	if err != nil {
		fmt.Println(err)
		c.JSON(404, gin.H{"error": "not found"})
	} else {
		c.JSON(200, categories)
	}
}

func CreateCategory(c *gin.Context) {
	var category model.Category

	c.BindJSON(&category)
	createdCategory, err := category.Save()
	if err != nil {
		fmt.Println(err)
		c.JSON(400, gin.H{"error": "unable to create category"})
	} else {
		c.JSON(200, createdCategory)
	}
}

func GetCategory(c *gin.Context) {
	id := c.Params.ByName("id")
	category, err := model.GetCategory(id)
	if err != nil {
		fmt.Println(err)
		c.JSON(404, gin.H{"error": "not found"})
	} else {
		c.JSON(200, category)
	}
}
