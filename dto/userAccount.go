package dto

type UserAccountRequest struct {
	FullName     string `json:"fullname" binding:"required"`
	Email        string `json:"email" binding:"required"`
	Password     string `json:"password" binding:"required"`
	ProfileImage string `json:"image"`
}

type UserAccountResponse struct {
	Id           uint   `json:"id"`
	Fullname     string `json:"fullname"`
	Email        string `json:"email"`
	Password     string `json:"password"`
	ProfileImage string `json:"image"`
}
