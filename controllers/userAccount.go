package controllers

import (
	"net/http"
	"plugspot/db"
	"plugspot/dto"
	"plugspot/model"

	"github.com/gin-gonic/gin"
)

type UserAccount struct{}

func (userAccount UserAccount) GetAllAccount(ctx *gin.Context) {
	var account []model.UserAccount
	db.Connection.Find(&account)

	var result []dto.UserAccountResponse
	for _, accounts := range account {
		result = append(result, dto.UserAccountResponse{
			Id:           accounts.ID,
			Fullname:     accounts.Fullname,
			Email:        accounts.Email,
			Password:     accounts.Password,
			ProfileImage: accounts.ProfileImage,
		})
	}

	ctx.JSON(http.StatusOK, result)
}

func (userAccount UserAccount) GetOneAccount(ctx *gin.Context) {
	id := ctx.Param("id")
	ctx.JSON(http.StatusOK, gin.H{"AccountId": id})
}

func (UserAccount UserAccount) CreateAccount(ctx *gin.Context) {
	var account dto.UserAccountRequest
	if err := ctx.ShouldBindJSON(&account); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userAccount := model.UserAccount{
		Fullname: account.FullName,
		Email:    account.Email,
		Password: account.Password,
	}

	if err := db.Connection.Create(&userAccount).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, dto.UserAccountResponse{
		Id:       userAccount.ID,
		Fullname: userAccount.Fullname,
		Email:    userAccount.Email,
		Password: userAccount.Password,
	})
}
