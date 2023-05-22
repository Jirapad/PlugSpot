package controllers

import (
	//"errors"
	"net/http"
	"os"
	"plugspot/db"
	"plugspot/dto"
	"plugspot/model"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	//"gorm.io/gorm"
)

type UserAccount struct{}

// func (userAccount UserAccount) GetAllAccount(ctx *gin.Context) {
// 	var account []model.UserAccount
// 	db.Connection.Find(&account)
// 	var result []dto.UserAccountResponse
// 	for _, accounts := range account {
// 		encryptedPassword, _ := bcrypt.GenerateFromPassword([]byte(accounts.Password), 10)
// 		result = append(result, dto.UserAccountResponse{
// 			Id:           accounts.ID,
// 			Fullname:     accounts.Fullname,
// 			Email:        accounts.Email,
// 			Password:     string(encryptedPassword),
// 			ProfileImage: accounts.ProfileImage,
// 		})
// 	}
// 	ctx.JSON(http.StatusOK, result)
// }

func (user UserAccount) Signup(ctx *gin.Context) {
	var account dto.UserAccountRequest
	if err := ctx.ShouldBindJSON(&account); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "failed to read body"})
		return
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(account.Password), 10)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "failed to hash password"})
	}
	userAccount := model.UserAccount{
		Fullname:    account.FullName,
		Email:       account.Email,
		PhoneNumber: "0" + strconv.Itoa(account.PhoneNumber),
		Role:        account.Role,
		Password:    string(hash),
	}
	if err := db.Connection.Create(&userAccount).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "failed to create account"})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"message": "created account"})
}

func (user UserAccount) Login(ctx *gin.Context) {
	var account dto.UserAccountLoginRequest
	if err := ctx.ShouldBindJSON(&account); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "failed to read body"})
		return
	}
	var userAccount model.UserAccount
	db.Connection.First(&userAccount, "email = ?", account.Email)
	if userAccount.ID == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid email"})
		return
	}
	err := bcrypt.CompareHashAndPassword([]byte(userAccount.Password), []byte(account.Password))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid password"})
		return
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userid": userAccount.ID,
		"role":   userAccount.Role,
		"exp":    time.Now().Add(time.Hour * 12).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "failed to create token"})
		return
	}
	ctx.SetSameSite(http.SameSiteLaxMode)
	ctx.SetCookie("Authorization", tokenString, 3600*12, "", "", false, true)
	ctx.JSON(http.StatusOK, gin.H{"role": userAccount.Role})
}

func (user UserAccount) CurrentUser(ctx *gin.Context) {
	account, _ := ctx.Get("user")
	ctx.JSON(http.StatusOK, gin.H{
		"message": account,
	})
}

func (user UserAccount) Update(ctx *gin.Context) {
	var account dto.UserAccountUpdateRequest
	if err := ctx.ShouldBind(&account); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	currentAccount, _ := ctx.Get("user")
	if currentAccount.(model.UserAccount).ID != account.UserId {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "can not edit other account"})
		return
	}
	var userAccount model.UserAccount
	// if err := db.Connection.First(&userAccount, account.UserId).Error; errors.Is(err, gorm.ErrRecordNotFound) {
	// 	ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
	// 	return
	// }
	db.Connection.First(&userAccount, account.UserId)
	if account.FullName != nil {
		db.Connection.Model(&userAccount).UpdateColumns(model.UserAccount{Fullname: *account.FullName})
		ctx.JSON(http.StatusOK, gin.H{"message": "update fullname success"})
	}
	if account.PhoneNumber != nil {
		db.Connection.Model(&userAccount).UpdateColumns(model.UserAccount{PhoneNumber: "0" + strconv.Itoa(*account.PhoneNumber)})
		ctx.JSON(http.StatusOK, gin.H{"message": "update phone number success"})
	}
	profileImage, _ := ctx.FormFile("profileImage")
	if profileImage != nil {
		profilePath := "./upload/userProfiles/" + uuid.New().String()
		ctx.SaveUploadedFile(profileImage, profilePath)
		db.Connection.Model(&userAccount).UpdateColumns(model.UserAccount{ProfileImage: profilePath})
		ctx.JSON(http.StatusOK, gin.H{"message": "update profile success"})
	}
}

func (user UserAccount) ResetPassword(ctx *gin.Context) {
	var account dto.UserAccountResetPasswordRequest
	if err := ctx.ShouldBindJSON(&account); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "failed to read body"})
		return
	}
	currentAccount, _ := ctx.Get("user")
	if currentAccount.(model.UserAccount).ID != account.UserId {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "can not edit other account"})
		return
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(account.NewPassword), 10)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "failed to hash password"})
	}
	var userAccount model.UserAccount
	db.Connection.First(&userAccount, account.UserId)
	db.Connection.Model(&userAccount).UpdateColumns(model.UserAccount{Password: string(hash)})
	ctx.JSON(http.StatusOK, gin.H{"message": "reset password success"})
}

func (user UserAccount) DeleteAccount(ctx *gin.Context) {
	var account dto.UserAccountDeleteAccountRequest
	if err := ctx.ShouldBindJSON(&account); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "failed to read body"})
		return
	}
	currentAccount, _ := ctx.Get("user")
	if currentAccount.(model.UserAccount).ID != account.UserId {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "can not edit other account"})
		return
	}
	db.Connection.Unscoped().Delete(&model.UserAccount{}, account.UserId)
	ctx.JSON(http.StatusOK, gin.H{"message": "delete account success"})
}
