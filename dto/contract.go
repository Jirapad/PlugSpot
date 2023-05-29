package dto

import "time"

type CreateContractRequest struct{
	CustomerId uint `json:"customerId" binding:"required"`
	ProviderId uint `json:"providerId" binding:"required"`
	StationId uint `json:"stationId" binding:"required"`
	CarId uint `json:"carId" binding:"required"`
	TimeSlot int `json:"timeSlot" binding:"required"`
}

type GetAllContractResponse struct{
	ContractId uint `json:"contractId"`
	StationName string `json:"stationName"`
	CustomeName string `json:"customerName"`
	ProviderName string `json:"providerName"`
	Date time.Time `json:"date"`
	TimeSlot int `json:"timeSlot"`
	Status string `json:"status"`
	TotalPrice float64 `json:"totalPrice"`
	PaymentMethod string`json:"paymentMethod"`
	CarPlate string `json:"carPlate"`
	CustomerId uint `json:"customerId"`
	ProviderId uint `json:"providerId"`
}

type UpdateContractStatusRequest struct{
	ProviderId uint `json:"providerId"`
	ContractId uint `json:"contractId" binding:"required"`
	TotalPrice *float64 `json:"totalPrice"`
	PaymentMethod *string `json:"paymentMedthod"`
	Status string `json:"status" binding:"required"`
}

type DeleteContractRequest struct{
	CustomerId uint `json:"customerId" binding:"required"`
	ContractId uint `json:"contractId" binding:"required"`
}