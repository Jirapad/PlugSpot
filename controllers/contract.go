package controllers

import (
	"net/http"
	"plugspot/db"
	"plugspot/dto"
	"plugspot/model"

	"github.com/gin-gonic/gin"
)

type Contract struct{}

func (con Contract) GetUserContract (ctx *gin.Context){
	var contract []model.Contract
	db.Connection.Find(&contract)
	currentAccount, _ := ctx.Get("user")
	var userContractList []dto.GetAllContractResponse
	for _,userContract := range contract{
		if userContract.CustomerId == currentAccount.(model.UserAccount).ID || userContract.ProviderId == currentAccount.(model.UserAccount).ID{
			var stationName model.Station
			db.Connection.First(&stationName,userContract.StationId)
			var customerName model.UserAccount
			db.Connection.First(&customerName,userContract.CustomerId)
			var providerName model.UserAccount
			db.Connection.First(&providerName,userContract.ProviderId)
			userContractList = append(userContractList, dto.GetAllContractResponse{
				ContractId: userContract.ID,
				StationName: stationName.StationName,
				CustomeName: customerName.Fullname,
				ProviderName: providerName.Fullname,
				Date: userContract.CreatedAt,
				TimeSlot: userContract.TimeSlot,
				Status: userContract.Status,
				TotalPrice: userContract.TotalPrice,
				PaymentMethod: userContract.PaymentMethod,
			})
		}
	}
	ctx.JSON(http.StatusOK, userContractList)
}
func (con Contract) CreateContract (ctx *gin.Context){
	var contract dto.CreateContractRequest
	if err := ctx.ShouldBindJSON(&contract); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "failed to read body"})
		return
	}
	currentAccount, _ := ctx.Get("user")
	if currentAccount.(model.UserAccount).ID != contract.CustomerId {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "can not edit other account"})
		return
	}
	//check car id have in database or not and match with customer id or not and car status is free or not 
	var allCars []model.Car
	db.Connection.Find(&allCars)
	for _,car := range allCars{
		if car.ID == contract.CarId && car.UserId == contract.CustomerId && car.CarStatus == "free"{
			//check station have in database or not and match with provider id or not
			var allStations []model.Station
			db.Connection.Find(&allStations)
			for _,station := range allStations{
				if station.ID == contract.StationId && station.UserId == contract.ProviderId{
					ContractInfo := model.Contract{
						CustomerId: currentAccount.(model.UserAccount).ID,
						ProviderId: contract.ProviderId,
						StationId: contract.StationId,
						CarId: contract.CarId,
						TimeSlot: contract.TimeSlot,
						Status: "in queue",
					}
					if err := db.Connection.Create(&ContractInfo).Error; err != nil {
						ctx.JSON(http.StatusBadRequest, gin.H{"error": "failed to create contract"})
						return
					}
					ctx.JSON(http.StatusCreated, gin.H{"message": "create contract success"})
					return
				}
			}
			ctx.JSON(http.StatusBadRequest, gin.H{"error":"not have station or provider id not match with station provider id"})
			return
		}		
	}
	ctx.JSON(http.StatusBadRequest,gin.H{
		"error":"not have customer car id or customer car id not match with customer id or car status is not free",
	})
}
func (con Contract) Update (ctx *gin.Context){
	var contract dto.UpdateContractStatusRequest
	if err := ctx.ShouldBindJSON(&contract); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "failed to read body"})
		return
	}
	currentAccount, _ := ctx.Get("user")
	if currentAccount.(model.UserAccount).ID != contract.ProviderId {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "can not edit other account"})
		return
	}
	var userContract model.Contract
	db.Connection.First(&userContract,contract.ContractId)
	if userContract.ID == 0 || userContract.ProviderId != contract.ProviderId{
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "don't have contract"})
		return
	}
	db.Connection.Model(&userContract).UpdateColumns(model.Contract{Status: contract.Status})
	ctx.JSON(http.StatusOK, gin.H{"message": "update contract status success"})
	if userContract.Status == "complete"{
		if contract.TotalPrice != nil{
			db.Connection.Model(&userContract).UpdateColumns(model.Contract{TotalPrice: *contract.TotalPrice})
			ctx.JSON(http.StatusOK, gin.H{"message": "update contract total price success"})
		}
		if contract.PaymentMethod != nil{
			db.Connection.Model(&userContract).UpdateColumns(model.Contract{PaymentMethod: *contract.PaymentMethod})
			ctx.JSON(http.StatusOK, gin.H{"message": "update contract payment method success"})
		}
	}
}
func (con Contract) DeleteContract (ctx *gin.Context){
	var contract dto.DeleteContractRequest
	if err := ctx.ShouldBindJSON(&contract); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "failed to read body"})
		return
	}
	currentAccount, _ := ctx.Get("user")
	if currentAccount.(model.UserAccount).ID != contract.CustomerId {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "can not edit other account"})
		return
	}
	var userContract model.Contract
	db.Connection.First(&userContract,contract.ContractId)
	if userContract.ID == 0 || userContract.CustomerId != contract.CustomerId{
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "don't have contract"})
		return
	}
	db.Connection.Unscoped().Delete(&model.Contract{}, contract.ContractId)
	ctx.JSON(http.StatusOK, gin.H{"message": "delete contract success"})
}