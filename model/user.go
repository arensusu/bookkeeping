package model

import (
	"firstapp/database"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"size:255;not null;unique" json:"username"`
	Password string `gorm:"size:255;not null" json:"-"`
	Admin    bool   `gorm:"default:false" json:"admin"`
}

func (user *User) Save() (*User, error) {
	err := database.Database.Create(&user).Error
	if err != nil {
		return &User{}, err
	}
	return user, nil
}

func GetAllUsers() (*[]User, error) {
	var users []User
	err := database.Database.Find(&users).Error
	if err != nil {
		return &[]User{}, err
	}
	return &users, nil
}

func GetUserById(id string) (User, error) {
	var user User
	err := database.Database.Where("ID=?", id).First(&user).Error
	if err != nil {
		return User{}, err
	} else {
		return user, nil
	}
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
