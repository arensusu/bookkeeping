package controller

import (
	"firstapp/model"
	"fmt"

	"github.com/gin-gonic/gin"
)

func GetAllCategories(c *gin.Context) {
	categories, err := model.GetAllCategories()
	if err != nil {
		c.JSON(400, gin.H{"error": "not found"})
		return
	}
	c.JSON(200, categories)
	return
}

func CreateCategory(c *gin.Context) {
	var category model.Category

	if err := c.BindJSON(&category); err != nil {
		fmt.Println(err)
		return
	}
	err := model.CreateCategory(&category)
	if err != nil {
		c.JSON(400, gin.H{"error": "unable to create category"})
		return
	}
	c.JSON(200, category)
	return
}

func GetCategory(c *gin.Context) {
	id := c.Params.ByName("id")
	category, err := model.GetCategoryByID(id)
	if err != nil {
		c.JSON(400, gin.H{"error": "not found"})
		return
	}
	c.JSON(200, category)
	return
}

func DeleteCategory(c *gin.Context) {
	id := c.Params.ByName("id")
	err := model.DeleteCategory(id)
	if err != nil {
		c.JSON(400, gin.H{"error": "not found"})
		return
	}
	c.JSON(200, gin.H{"message": "category deleted"})
	return
}

func UpdateCategory(c *gin.Context) {
	id := c.Params.ByName("id")
	var category model.Category
	if err := c.BindJSON(&category); err != nil {
		fmt.Println(err)
		return
	}
	err := model.UpdateCategory(id, &category)
	if err != nil {
		fmt.Println(err)
		c.JSON(400, gin.H{"error": "unable to update category"})
		return
	}
	c.JSON(200, gin.H{"message": "category updated"})
	return

}
