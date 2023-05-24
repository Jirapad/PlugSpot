package main

import (
	"log"
	"os"
	"plugspot/db"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	if os.Getenv("APP_ENV") == "production" {
		gin.SetMode(gin.ReleaseMode)
	} else {
		if err := godotenv.Load(); err != nil {
			log.Fatal("Error loading .env file")
		}
	}
	db.ConnectDB()
	db.Migrate()
	os.MkdirAll("upload/userProfiles",0755)
	os.MkdirAll("upload/stations",0755)
	r := gin.Default()
	serverRoutes(r)
	r.Run(":1234")
	//r.Run()
}
