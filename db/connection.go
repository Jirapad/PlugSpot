package db

import (
	"log"
	"os"
	"plugspot/model"

	//"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var Connection *gorm.DB

func ConnectDB() {
	dsn := os.Getenv("DATABASE_DSN")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{Logger: logger.Default.LogMode(logger.Info)})
	if err != nil {
		log.Fatal("Can not connect to the database")
	}

	Connection = db
}

func Migrate() {
	Connection.AutoMigrate(
		&model.UserAccount{},
		&model.Car{},
		&model.Station{},
		&model.TimeSlot{},
		&model.Contract{},
	)
}
