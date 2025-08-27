package models

import (
	"time"

	"gorm.io/gorm"
)

type AuditTrailModel struct {
	gorm.Model
	UserID    *uint     `json:"user_id"`
	Entity    string    `json:"entity"`    // e.g. "devices", "users", "device_locations"
	EntityID  uint      `json:"entity_id"` // row id of the entity
	Action    string    `json:"action"`    // create, update, delete, login, etc.
	OldValue  string    `json:"old_value"`
	NewValue  string    `json:"new_value"`
	IPAddress string    `json:"ip_address"`
	Timestamp time.Time `json:"timestamp"`
}

func (AuditTrailModel) TableName() string {
	return "audit_trails"
}
