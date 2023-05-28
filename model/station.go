package model

import "gorm.io/gorm"

type Station struct {
	gorm.Model
	UserId        uint         `gorm:"uniqueIndex;not null"`
	StationName   string       `gorm:"type:varchar(150);not null"`
	StationImage  string       `gorm:"uniqueIndex;type:varchar(150);not null"`
	StationDetail string       `gorm:"type:varchar(150);not null"`
	ProviderPhone string       `gorm:"type:varchar(150);not null"`
	Latitude      string       `gorm:"uniqueIndex;type:varchar(150);not null"`
	Longitude     string       `gorm:"uniqueIndex;type:varchar(150);not null"`
	Timeslots     []TimeSlot `gorm:"foreignKey:StationId"`
}
