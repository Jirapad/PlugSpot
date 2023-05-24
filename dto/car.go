package dto

type AddNewCarRequest struct {
	UserId   uint   `json:"userId" binding:"required"`
	CarPlate string `json:"carPlate" binding:"required"`
	CarBrand string `json:"carBrand" binding:"required"`
	CarModel string `json:"carModel" binding:"required"`
}

type DeleteUserCarRequest struct {
	UserId   uint   `json:"userId" binding:"required"`
	CarPlate string `json:"carPlate" binding:"required"`
}

type CarsResponse struct {
	CarId     uint   `json:"carId"`
	CarPlate  string `json:"carPlate"`
	CarBrand  string `json:"carBrand"`
	CarModel  string `json:"carModel"`
	CarStatus string `json:"carStatus"`
}

type CarsCheckResponse struct {
	CarId    uint   `json:"carId"`
	CarPlate string `json:"carPlate"`
}

type CarUpdateRequest struct {
	UserId    uint   `json:"userId" binding:"required"`
	CarId     uint   `json:"carId" binding:"required"`
	CarPlate  *string `json:"carPlate"`
	CarBrand  *string `json:"carBrand"`
	CarModel  *string `json:"carModel"`
	CarStatus *string `json:"carStatus"`
}
