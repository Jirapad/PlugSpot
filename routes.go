package main

import (
	"plugspot/controllers"

	"github.com/gin-gonic/gin"
)

func serverRoutes(r *gin.Engine) {

	//UserAccount
	userAccountController := controllers.UserAccount{}
	userAccount := r.Group("/userAccounts")
	userAccount.GET("", userAccountController.GetAllAccount)
	userAccount.GET("/:id", userAccountController.GetOneAccount)
	userAccount.POST("", userAccountController.CreateAccount)
	userAccount.PATCH("")
	userAccount.DELETE("")
}
