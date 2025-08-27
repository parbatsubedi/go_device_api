package models

import (
	"time"

	"gorm.io/gorm"
)

type DeviceLocationModel struct {
	gorm.Model
	DeviceID     uint      `json:"device_id"`
	Latitude     float64   `json:"latitude"`
	Longitude    float64   `json:"longitude"`
	Accuracy     float64   `json:"accuracy"`
	Altitude     float64   `json:"altitude"`
	Speed        float64   `json:"speed"`
	BatteryLevel int       `json:"battery_level"`
	NetworkType  string    `json:"network_type"` // wifi, gsm, gps, lte, 5g
	IsCharging   bool      `json:"is_charging"`
	IPAddress    string    `json:"ip_address"`
	RecordedAt   time.Time `json:"recorded_at"`
}

func (DeviceLocationModel) TableName() string {
	return "device_locations"
}
