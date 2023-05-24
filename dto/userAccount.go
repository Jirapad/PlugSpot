package dto

type UserAccountRequest struct {
	FullName     string `json:"fullname" binding:"required"`
	Email        string `json:"email" binding:"required"`
	Password     string `json:"password" binding:"required"`
	PhoneNumber  string    `json:"phonenumber" binding:"required"`
	Role         string `json:"role" binding:"required"`
	ProfileImage string `json:"image"`
}

type UserAccountLoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserAccountUpdateRequest struct {
	UserId      uint    `form:"userId" binding:"required"`
	FullName    *string `form:"fullname"`
	PhoneNumber *int    `form:"phoneNumber"`
}

type UserAccountResetPasswordRequest struct {
	UserId      uint   `json:"userId" binding:"required"`
	OldPassword string `json:"oldPassword" binding:"required"`
	NewPassword string `json:"newPassword" binding:"required"`
}

type UserAccountDeleteAccountRequest struct {
	UserId uint `json:"userId" binding:"required"`
}
