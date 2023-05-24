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
	userAccount.GET("/currentuser", middleware.MiddlewareForAllRole, userAccountController.CurrentUser)
	userAccount.POST("/signup", userAccountController.Signup)
	userAccount.POST("/login", userAccountController.Login)
	userAccount.PATCH("/update", middleware.MiddlewareForAllRole, userAccountController.Update)
	userAccount.PATCH("/resetpassword", middleware.MiddlewareForAllRole, userAccountController.ResetPassword)
	userAccount.DELETE("/deleteaccount", middleware.MiddlewareForAllRole, userAccountController.DeleteAccount)

	//Car
	carController := controllers.Car{}
	car := r.Group("/car")
	car.GET("/getallusercars",middleware.MiddlewareForCustomerRole,carController.GetAllUserCars)
	car.POST("/addnewcar",middleware.MiddlewareForCustomerRole,carController.AddNewCar)
	car.PATCH("/update",middleware.MiddlewareForCustomerRole,carController.Update)
	car.DELETE("/deleteusercar",middleware.MiddlewareForCustomerRole,carController.DeleteUserCar)
}
