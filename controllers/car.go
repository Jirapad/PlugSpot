package controllers

// import (
// 	"net/http"
// 	"plugspot/db"
// 	"plugspot/dto"
// 	"plugspot/model"

// 	"github.com/gin-gonic/gin"
// )

// type Car struct{}

// func (carInfo Car) GetAllCar(ctx *gin.Context) {
// 	var car []model.Car
// 	db.Connection.Find(&car)

// 	var result []dto.CarResponse
// 	for _, cars := range car {
// 		result = append(result, dto.CarResponse{
// 			Id:           cars.ID,
// 			CarName:     cars.CarName,
// 			ChargeType:        cars.ChargeType,
// 			UserId:     cars.UserId,
// 		})
// 	}
// 	ctx.JSON(http.StatusOK, result)
// }
