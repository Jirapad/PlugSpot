package controllers

import (
	"net/http"
	"plugspot/db"
	"plugspot/dto"
	"plugspot/model"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type Station struct{}

func (location Station) GetUserStation(ctx *gin.Context) {
	var allStation []model.Station
	db.Connection.Find(&allStation)
	currentAccount, _ := ctx.Get("user")
	for _, station := range allStation {
		if currentAccount.(model.UserAccount).ID == station.UserId {
			var allTimeSlots []model.TimeSlot
			db.Connection.Find(&allTimeSlots)
			var stationTimeSlot []dto.StationTimeSlotsResponse
			for _, timeSlot := range allTimeSlots {
				if station.ID == timeSlot.StationId {
					stationTimeSlot = append(stationTimeSlot, dto.StationTimeSlotsResponse{
						TimeSlotNo: timeSlot.TimeSlotNo,
						Status:     timeSlot.Status,
					})
				}
			}
			UserStation := dto.StationResponse{
				StationId:     station.ID,
				StationName:   station.StationName,
				StationImage:  station.StationImage,
				StationDetail: station.StationDetail,
				ProviderPhone: station.ProviderPhone,
				Latitude:      station.Latitude,
				Longitude:     station.Longitude,
				Timeslots:     stationTimeSlot,
			}
			ctx.JSON(http.StatusOK, UserStation)
		}
	}
}

func (location Station) UpdateStation(ctx *gin.Context) {
	var station dto.UpdateStationRequest
	if err := ctx.ShouldBind(&station); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "failed to read body"})
		return
	}
	currentAccount, _ := ctx.Get("user")
	if currentAccount.(model.UserAccount).ID != station.UserId {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "can not edit other account"})
		return
	}
	var userStation model.Station
	db.Connection.First(&userStation, "user_id = ?", station.UserId)
	if userStation.ID == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "don't have station"})
		return
	}
	if station.StationName != nil {
		db.Connection.Model(&userStation).UpdateColumns(model.Station{StationName: *station.StationName})
		ctx.JSON(http.StatusOK, gin.H{"message": "update station name success"})
	}
	if station.StationDetail != nil {
		db.Connection.Model(&userStation).UpdateColumns(model.Station{StationDetail: *station.StationDetail})
		ctx.JSON(http.StatusOK, gin.H{"message": "update station detail success"})
	}
	if station.Latitude != nil {
		db.Connection.Model(&userStation).UpdateColumns(model.Station{Latitude: *station.Latitude})
		ctx.JSON(http.StatusOK, gin.H{"message": "update station latitude success"})
	}
	if station.Longitude != nil {
		db.Connection.Model(&userStation).UpdateColumns(model.Station{Longitude: *station.Longitude})
		ctx.JSON(http.StatusOK, gin.H{"message": "update station longitude success"})
	}
	stationImage, _ := ctx.FormFile("stationImage")
	if stationImage != nil {
		stationPath := "./upload/stations/" + uuid.New().String()
		ctx.SaveUploadedFile(stationImage, stationPath)
		db.Connection.Model(&userStation).UpdateColumns(model.Station{StationImage: stationPath})
		ctx.JSON(http.StatusOK, gin.H{"message": "update station image success"})
	}
}

func (location Station) AddNewStation(ctx *gin.Context) {
	var station dto.AddNewStationRequest
	if err := ctx.ShouldBind(&station); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "failed to read body"})
		return
	}
	currentAccount, _ := ctx.Get("user")
	if currentAccount.(model.UserAccount).ID != station.UserId {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "can not edit other account"})
		return
	}

	stationImage, _ := ctx.FormFile("stationImage")
	stationPath := "./upload/stations/" + uuid.New().String()
	ctx.SaveUploadedFile(stationImage, stationPath)

	var createTimeSlot []model.TimeSlot
	for time := 0; time < 12; time++ {
		createTimeSlot = append(createTimeSlot, model.TimeSlot{
			TimeSlotNo: time + 1,
			Status:     "free",
		})
	}
	stationInformation := model.Station{
		UserId:        station.UserId,
		StationName:   station.StationName,
		StationImage:  stationPath,
		StationDetail: station.StationDetail,
		ProviderPhone: currentAccount.(model.UserAccount).PhoneNumber,
		Latitude:      station.Latitude,
		Longitude:     station.Longitude,
		Timeslots:     createTimeSlot,
	}
	if err := db.Connection.Create(&stationInformation).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "failed to add new station"})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"message": "add new station success"})
}

func (location Station) DeleteStation(ctx *gin.Context) {
	var station dto.DeleteUserStationRequest
	if err := ctx.ShouldBindJSON(&station); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "failed to read body"})
		return
	}
	currentAccount, _ := ctx.Get("user")
	if currentAccount.(model.UserAccount).ID != station.UserId {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "can not edit other account"})
		return
	}
	var allStations []model.Station
	db.Connection.Find(&allStations)
	for _, userStation := range allStations {
		if userStation.UserId == station.UserId && userStation.ID == station.StationId {
			var allTimeSlots []model.TimeSlot
			db.Connection.Find(&allTimeSlots)
			for _, timeSlot := range allTimeSlots {
				if timeSlot.StationId == station.StationId {
					db.Connection.Unscoped().Delete(&model.TimeSlot{}, timeSlot.ID)
				}
			}
			ctx.JSON(http.StatusOK, gin.H{"message": "delete time slots success"})
			db.Connection.Unscoped().Delete(&model.Station{}, station.StationId)
			ctx.JSON(http.StatusOK, gin.H{"message": "delete station success"})
			return
		}
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "no station id"})
}

func (location Station) TimeSlotUpdate(ctx *gin.Context) {
	var timeSlot dto.TimeSlotUpdateRequest
	if err := ctx.ShouldBindJSON(&timeSlot); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "failed to read body"})
		return
	}
	currentAccount, _ := ctx.Get("user")
	if currentAccount.(model.UserAccount).ID != timeSlot.UserId {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "can not edit other account"})
		return
	}
	var userStation model.Station
	db.Connection.First(&userStation, "user_id = ?", timeSlot.UserId)
	if userStation.ID == 0 || userStation.ID != timeSlot.StationId {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "don't have station"})
		return
	}
	var stationTimeSlot []model.TimeSlot
	db.Connection.Find(&stationTimeSlot)
	for _, slot := range stationTimeSlot {
		if slot.StationId == timeSlot.StationId && slot.TimeSlotNo == timeSlot.TimeSlotNo {
			db.Connection.Model(&slot).UpdateColumns(model.TimeSlot{Status: timeSlot.Status})
			ctx.JSON(http.StatusOK, gin.H{"message": "update station time slot success"})
			return
		}
	}
	ctx.JSON(http.StatusBadRequest, gin.H{"error": "don't have that time slot"})
}

func (location Station) GetAllStation(ctx *gin.Context){
	var allStation []model.Station
	db.Connection.Find(&allStation)
	var stationTimeSlot []model.TimeSlot
	db.Connection.Find(&stationTimeSlot)
	var result []model.Station
	for _,station := range allStation{
		var slot []model.TimeSlot
		for _,timeSlot := range stationTimeSlot{
			if timeSlot.StationId == station.ID{
				slot = append(slot, timeSlot)
			}
		}
		station.Timeslots = slot
		result = append(result, station) 
	}
	ctx.JSON(http.StatusOK, result)
}

func (location Station) ResetAllTimeSlot(ctx *gin.Context){
	
}