package models

import (
	"time"

	"gorm.io/gorm"
)

type DeviceLocationModel struct {
	gorm.Model
	DeviceID    uint      `json:"device_id"`
	Latitude    float64   `json:"latitude"`
	Longitude   float64   `json:"longitude"`
	Accuracy    float64   `json:"accuracy"`
	NetworkType string    `json:"network_type"` // wifi, gsm, gps
	RecordedAt  time.Time `json:"recorded_at"`
}

func (DeviceLocationModel) TableName() string {
	return "device_locations"
}
