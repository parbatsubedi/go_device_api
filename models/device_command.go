package models

import (
	"time"

	"gorm.io/gorm"
)

// DeviceCommandModel represents a command sent to a device
type DeviceCommandModel struct {
	gorm.Model
	DeviceID    uint   `json:"device_id" gorm:"not null;index"`
	CommandType string `json:"command_type" gorm:"type:varchar(255);not null"`
	CommandData string `json:"command_data" gorm:"type:text"`
	Status      string `json:"status" gorm:"type:varchar(100);default:'pending'"` // e.g., pending, sent, acknowledged, failed
	SentAt      *time.Time `json:"sent_at"`
	AckedAt     *time.Time `json:"acked_at"`
}

// TableName overrides the table name used by DeviceCommandModel to `device_commands`
func (DeviceCommandModel) TableName() string {
	return "device_commands"
}
