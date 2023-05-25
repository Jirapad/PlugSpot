package dto

type AddNewStationRequest struct {
	UserId        uint   `form:"userId" binding:"required"`
	StationName   string `form:"stationName" binding:"required"`
	StationDetail string `form:"stationDetail" binding:"required"`
	Latitude      string `form:"latitude" binding:"required"`
	Longitude     string `form:"longitude" binding:"required"`
}

type DeleteUserStationRequest struct {
	UserId    uint `json:"userId" binding:"required"`
	StationId uint `json:"stationId" binding:"required"`
}

type UpdateStationRequest struct {
	UserId        uint    `form:"userId" binding:"required"`
	StationName   *string `form:"stationName"`
	StationDetail *string `form:"stationDetail"`
	Latitude      *string `form:"latitude"`
	Longitude     *string `form:"longitude"`
}

type StationResponse struct {
	StationId     uint                       `json:"stationId"`
	StationName   string                     `json:"stationName"`
	StationImage  string                     `json:"stationImage"`
	StationDetail string                     `json:"stationDetail"`
	ProviderPhone string                     `json:"providerPhone"`
	Latitude      string                     `json:"latitude"`
	Longitude     string                     `json:"longitude"`
	Timeslots     []StationTimeSlotsResponse `json:"timeSlots"`
}

type StationTimeSlotsResponse struct {
	TimeSlotNo int    `json:"timeSlotNo"`
	Status     string `json:"status"`
}

type TimeSlotUpdateRequest struct {
	UserId     uint   `json:"userId" binding:"required"`
	StationId  uint   `json:"stationId" binding:"required"`
	TimeSlotNo int    `json:"timeSlotNo" binding:"required"`
	Status     string `json:"status" binding:"required"`
}
