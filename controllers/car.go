package controllers

import (
	"net/http"
	"plugspot/db"
	"plugspot/dto"
	"plugspot/model"

	"github.com/gin-gonic/gin"
)

type Car struct{}

func (carInfo Car) GetAllUserCars(ctx *gin.Context) {
	var allCars []model.Car
	db.Connection.Find(&allCars)
	currentAccount, _ := ctx.Get("user")
	var allUserCar []dto.CarsResponse
	for _, car := range allCars {
		if currentAccount.(model.UserAccount).ID == car.UserId {
			allUserCar = append(allUserCar, dto.CarsResponse{
				CarId:     car.ID,
				CarPlate:  car.CarPlate,
				CarBrand:  car.CarBrand,
				CarModel:  car.CarModel,
				CarStatus: car.CarStatus,
			})
		}
	}
	ctx.JSON(http.StatusOK, allUserCar)
}
func (carInfo Car) AddNewCar(ctx *gin.Context) {
	var car dto.AddNewCarRequest
	if err := ctx.ShouldBindJSON(&car); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "failed to read body"})
		return
	}
	currentAccount, _ := ctx.Get("user")
	if currentAccount.(model.UserAccount).ID != car.UserId {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "can not edit other account"})
		return
	}
	carInformation := model.Car{
		UserId:    car.UserId,
		CarPlate:  car.CarPlate,
		CarBrand:  car.CarBrand,
		CarModel:  car.CarModel,
		CarStatus: "free",
	}
	if err := db.Connection.Create(&carInformation).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "failed to add new car"})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"message": "add new car success"})
}
func (carInfo Car) Update(ctx *gin.Context) {
	var car dto.CarUpdateRequest
	if err := ctx.ShouldBindJSON(&car); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "failed to read body"})
		return
	}
	//check userId
	currentAccount, _ := ctx.Get("user")
	if currentAccount.(model.UserAccount).ID != car.UserId {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "can not edit other account"})
		return
	}
	//list of userCar
	var allCars []model.Car
	db.Connection.Find(&allCars)
	var allUserCar []dto.CarsCheckResponse
	for _, cars := range allCars {
		if currentAccount.(model.UserAccount).ID == cars.UserId {
			allUserCar = append(allUserCar, dto.CarsCheckResponse{
				CarId:    cars.ID,
				CarPlate: cars.CarPlate,
			})
		}
	}
	//check carId
	for _, carId := range allUserCar {
		if car.CarId == carId.CarId {
			var carUpdate model.Car
			db.Connection.First(&carUpdate, car.CarId)
			//update carPlate
			if car.CarPlate != nil {
				//check new carPlate is already exist or not
				for _, plate := range allCars {
					if plate.CarPlate == *car.CarPlate {
						ctx.JSON(http.StatusBadRequest, gin.H{"error": "car plate already exist"})
						return
					}
				}
				db.Connection.Model(&carUpdate).UpdateColumns(model.Car{CarPlate: *car.CarPlate})
				ctx.JSON(http.StatusOK, gin.H{"message": "update car plate success"})
			}
			if car.CarBrand != nil {
				db.Connection.Model(&carUpdate).UpdateColumns(model.Car{CarBrand: *car.CarBrand})
				ctx.JSON(http.StatusOK, gin.H{"message": "update car car brand success"})
			}
			if car.CarModel != nil {
				db.Connection.Model(&carUpdate).UpdateColumns(model.Car{CarModel: *car.CarModel})
				ctx.JSON(http.StatusOK, gin.H{"message": "update car model success"})
			}
			if car.CarStatus != nil {
				db.Connection.Model(&carUpdate).UpdateColumns(model.Car{CarStatus: *car.CarStatus})
				ctx.JSON(http.StatusOK, gin.H{"message": "update car status success"})
			}
			return
		}
	}
	ctx.JSON(http.StatusBadRequest, gin.H{"error": "can not find your car id"})
}
func (carInfo Car) DeleteUserCar(ctx *gin.Context) {
	var car dto.DeleteUserCarRequest
	if err := ctx.ShouldBindJSON(&car); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "failed to read body"})
		return
	}
	currentAccount, _ := ctx.Get("user")
	if currentAccount.(model.UserAccount).ID != car.UserId {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "can not edit other account"})
		return
	}
	var allCars []model.Car
	db.Connection.Find(&allCars)
	//var allUserCar []dto.CarsCheckResponse
	for _, cars := range allCars {
		if currentAccount.(model.UserAccount).ID == cars.UserId && cars.ID == car.CarId{
			db.Connection.Unscoped().Delete(&model.Car{}, car.CarId)
			ctx.JSON(http.StatusOK, gin.H{"message": "delete user car success"})
			return
			// allUserCar = append(allUserCar, dto.CarsCheckResponse{
			// 	CarId:    cars.ID,
			// 	CarPlate: cars.CarPlate,
			// })
		}
	}
	// for _, plate := range allUserCar {
	// 	if car.CarPlate == plate.CarPlate {
	// 		db.Connection.Unscoped().Delete(&model.Car{}, plate.CarId)
	// 		ctx.JSON(http.StatusOK, gin.H{"message": "delete user car success"})
	// 		return
	// 	}
	// }
	ctx.JSON(http.StatusBadRequest, gin.H{"error": "can not find your car"})
}
