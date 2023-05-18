package model

import (
	"firstapp/database"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"size:255;not null" json:"username"`
	Password string `gorm:"size:255;not null" json:"-"`
	Admin    bool   `gorm:"default:false" json:"admin"`
}

func CreateUser(user *User) error {
	if err := database.Database.Create(user).Error; err != nil {
		return err
	}
	return nil
}

func GetUserByName(username string) (User, error) {
	var user User
	err := database.Database.Where("username=?", username).First(&user).Error
	if err != nil {
		return User{}, err
	} else {
		return user, nil
	}
}

func DeleteUser(id uint) error {
	err := database.Database.Delete(&User{}, id).Error
	if err != nil {
		return err
	} else {
		return nil
	}
}

func UpdateUser(id uint, user *User) error {
	err := database.Database.Save(&user).Error
	if err != nil {
		return err
	} else {
		return nil
	}
}
