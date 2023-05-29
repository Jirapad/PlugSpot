package model

import "gorm.io/gorm"

type Contract struct{
	gorm.Model
	CustomerId uint
	ProviderId uint
	StationId uint 
	CarId uint
	TimeSlot int `gorm:"not null"`
	Status string `gorm:"type:varchar(150);not null"`
	TotalPrice float64 `gorm:"not null"`
	PaymentMethod string`gorm:"type:varchar(150);not null"`
}