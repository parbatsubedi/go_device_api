package apirequests

import "go_api/models"

type CreateDeviceLocationRequest struct {
	DeviceID  uint    `json:"device_id" binding:"required"`
	Latitude  float64 `json:"latitude" binding:"required,latitude"`
	Longitude float64 `json:"longitude" binding:"required,longitude"`
	Altitude  float64 `json:"altitude"`
	Speed     float64 `json:"speed"`
}

func (r *CreateDeviceLocationRequest) ToModel() models.DeviceLocationModel {
	return models.DeviceLocationModel{
		DeviceID:  r.DeviceID,
		Latitude:  r.Latitude,
		Longitude: r.Longitude,
	}
}
