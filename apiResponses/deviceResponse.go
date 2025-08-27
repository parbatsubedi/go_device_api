package apiresponses

import "go_api/models"

type DeviceResponse struct {
	SuccessResponse
	Data models.DeviceModelDto `json:"data"`
}

func NewDeviceResponse() DeviceResponse {
	return DeviceResponse{
		NewSuccessResponse(),
		models.DeviceModelDto{},
	}
}

type DeviceListResponse struct {
	SuccessResponse
	// Data models.Pagination `json:"data"`
	Data []models.DeviceModelDto `json:"data"`
}

func NewDeviceListResponse() DeviceListResponse {
	return DeviceListResponse{
		NewSuccessResponse(),
		// models.Pagination{},
		[]models.DeviceModelDto{},
	}
}
