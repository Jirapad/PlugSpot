package model

import "gorm.io/gorm"

type TimeSlot struct{
	gorm.Model
	StationId uint 
	TimeSlotNo int `gorm:"not null"`
	Status string `gorm:"type:varchar(50);not null"`
}