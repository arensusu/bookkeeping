package model

import (
	"bookkeeping-backend/database"

	"gorm.io/gorm"
)

type Category struct {
	gorm.Model `json:"-"`
	Name       string `gorm:"size:255;not null;unique" json:"name"`
}

func CreateCategory(category *Category) error {
	err := database.Database.Create(category).Error
	if err != nil {
		return err
	}
	return nil
}

func GetAllCategories() (*[]Category, error) {
	var categories []Category
	err := database.Database.Find(&categories).Error
	if err != nil {
		return nil, err
	}
	return &categories, nil
}

func GetCategoryByID(id string) (Category, error) {
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

func DeleteCategory(id string) error {
	err := database.Database.Delete(&Category{}, id).Error
	if err != nil {
		return err
	} else {
		return nil
	}
}

func UpdateCategory(id string, category *Category) error {
	err := database.Database.Save(category).Error
	if err != nil {
		return err
	} else {
		return nil
	}
}
