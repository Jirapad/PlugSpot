package model

import "gorm.io/gorm"

type UserAccount struct {
	gorm.Model
	Fullname     string `gorm:"type:varchar(200);not null"`
	Email        string `gorm:"type:varchar(150);not null"`
	Password     string `gorm:"type:varchar(150);not null"`
	ProfileImage string `gorm:"type:varchar(150);not null"`
}
