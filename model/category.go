package model

import (
	"firstapp/database"

	"gorm.io/gorm"
)

type Category struct {
	gorm.Model `json:"-"`
	Name       string `gorm:"size:255;not null;unique" json:"name"`
}

func (category *Category) Save() (*Category, error) {
	err := database.Database.Create(&category).Error
	if err != nil {
		return &Category{}, err
	}
	return category, nil
}

func GetAllCategories() (*[]Category, error) {
	var categories []Category
	err := database.Database.Find(&categories).Error
	if err != nil {
		return &[]Category{}, err
	}
	return &categories, nil
}

func GetCategory(id string) (Category, error) {
	var category Category
	err := database.Database.Where("ID=?", id).First(&category).Error
	if err != nil {
		return Category{}, err
	} else {
		return category, nil
	}
}

func GetCategoryByName(name string) (Category, error) {
	var category Category
	err := database.Database.Where("name=?", name).First(&category).Error
	if err != nil {
		return Category{}, err
	} else {
		return category, nil
	}
}
