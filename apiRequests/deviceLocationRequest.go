package apirequests

import (
	"go_api/models"
	"time"
)

type CreateDeviceLocationRequest struct {
	DeviceID     uint    `json:"device_id" binding:"required"`
	Latitude     float64 `json:"latitude" binding:"required,latitude"`
	Longitude    float64 `json:"longitude" binding:"required,longitude"`
	Accuracy     float64 `json:"accuracy"`
	Altitude     float64 `json:"altitude"`
	Speed        float64 `json:"speed"`
	BatteryLevel int     `json:"battery_level"`
	NetworkType  string  `json:"network_type"` // wifi, gsm, gps, lte, 5g
	IsCharging   bool    `json:"is_charging"`
	IPAddress    string  `json:"ip_address"`
}

func (r *CreateDeviceLocationRequest) ToModel() models.DeviceLocationModel {
	return models.DeviceLocationModel{
		DeviceID:     r.DeviceID,
		Latitude:     r.Latitude,
		Longitude:    r.Longitude,
		Accuracy:     r.Accuracy,
		Altitude:     r.Altitude,
		Speed:        r.Speed,
		BatteryLevel: r.BatteryLevel,
		NetworkType:  r.NetworkType,
		IsCharging:   r.IsCharging,
		IPAddress:    r.IPAddress,
		RecordedAt:   time.Now(),
	}
}
