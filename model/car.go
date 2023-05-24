package model

import "gorm.io/gorm"

type Car struct {
	gorm.Model
	UserId    uint   `gorm:"not null"`
	CarPlate  string `gorm:"uniqueIndex;type:varchar(150);not null"`
	CarBrand  string `gorm:"type:varchar(200);not null"`
	CarModel  string `gorm:"type:varchar(200);not null"`
	CarStatus string `gorm:"type:varchar(50);not null"`
}
