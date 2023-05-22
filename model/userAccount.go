package model

import "gorm.io/gorm"

type UserAccount struct {
	gorm.Model
	Fullname     string `gorm:"type:varchar(200);not null"`
	Email        string `gorm:"uniqueIndex;type:varchar(150);not null"`
	Password     string `gorm:"type:varchar(150);not null"`
	PhoneNumber  string `gorm:"type:varchar(10);not null"`
	Role         string `grom:"uniqueIndex;type:varchar(50);notnull"`
	ProfileImage string `gorm:"type:varchar(150);not null"`
}
