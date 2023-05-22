package model

import (
	"bookkeeping-backend/database"

	"gorm.io/gorm"
)

type Detail struct {
	gorm.Model
	ID         uint     `gorm:"primaryKey;autoIncrement" json:"id"`
	UserID     uint     `gorm:"not null" json:"-"`
	User       User     `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"user"`
	CategoryID uint     `gorm:"not null" json:"-"`
	Category   Category `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"category"`
	Cost       int      `gorm:"not null" json:"cost"`
	Date       string   `gorm:"size:255;not null" json:"date"`
}

func CreateDetail(detail *Detail) error {
	err := database.Database.Create(detail).Error
	if err != nil {
		return err
	}
	return nil
}

func GetAllDetailsOfUser(userID uint, startDate string, endDate string) (*[]Detail, error) {
	var details []Detail
	data := database.Database.Preload("User").Preload("Category").Where("user_id=?", userID)
	if startDate != "" {
		data = data.Where("created_at>=?", startDate)
	}
	if endDate != "" {
		data = data.Where("created_at<=?", endDate)
	}

	if err := data.Find(&details).Error; err != nil {
		return &[]Detail{}, err
	}
	return &details, nil
}

func GetDetail(id string) (Detail, error) {
	var detail Detail
	err := database.Database.Preload("User").Preload("Category").Where("id=?", id).First(&detail).Error
	if err != nil {
		return Detail{}, err
	} else {
		return detail, nil
	}
}

func DeleteDetail(id string) error {
	err := database.Database.Delete(&Detail{}, id).Error
	if err != nil {
		return err
	} else {
		return nil
	}
}
