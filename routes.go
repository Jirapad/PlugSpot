package main

import (
	"plugspot/controllers"
	"plugspot/middleware"

	"github.com/gin-gonic/gin"
)

func serverRoutes(r *gin.Engine) {

	//UserAccount
	userAccountController := controllers.UserAccount{}
	userAccount := r.Group("/userAccount")
	userAccount.POST("/signup", userAccountController.Signup)
	userAccount.POST("/login", userAccountController.Login)
	userAccount.PATCH("/update", middleware.MiddlewareForAllRole, userAccountController.Update)
	userAccount.PATCH("/resetpassword", middleware.MiddlewareForAllRole, userAccountController.ResetPassword)
	userAccount.GET("/currentuser", middleware.MiddlewareForAllRole, userAccountController.CurrentUser)

	//Car
	// carController := controllers.Car{}
	// car := r.Group("/cars")
	// car.GET("/getAllCars",carController.GetAllCar)
}
