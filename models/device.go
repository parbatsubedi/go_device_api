package models

import (
	"time"

	"github.com/jinzhu/copier"
	"gorm.io/gorm"
)

type DeviceModel struct {
	gorm.Model
	DeviceName   string    `json:"device_name" gorm:"type:varchar(255);not null"`
	DeviceIMEI1  string    `json:"device_imei1" gorm:"type:varchar(255);not null;uniqueIndex"`
	DeviceIMEI2  string    `json:"device_imei2" gorm:"type:varchar(255);uniqueIndex"`
	Manufacturer string    `json:"manufacturer" gorm:"type:varchar(255); not null"`
	DeviceModel  string    `json:"device_model" gorm:"type:varchar(255); not null"`
	Status       bool      `json:"status" gorm:"default:true"`
	LastSeenAt   time.Time `json:"last_seen_at"`
	UserID       uint      `json:"user_id"`

	User      UserModel           `gorm:"foreignKey:UserID"`
	Locations []DeviceLocationModel `gorm:"foreignKey:DeviceID"` // has many relation with DeviceLocationModel
}

func (DeviceModel) TableName() string {
	return "devices"
}

type DeviceModelDto struct {
	ID           uint    `json:"ID"`
	DeviceName   string  `json:"device_name"`
	DeviceIMEI1  string  `json:"device_imei1"`
	DeviceIMEI2  string  `json:"device_imei2"`
	Manufacturer string  `json:"manufacturer"`
	DeviceModel  string  `json:"device_model"`
	Status       bool    `json:"status"`
	CreatedAt    string  `json:"created_at"`
	UpdatedAt    string  `json:"updated_at"`
	DeletedAt    *string `json:"deleted_at"` // * indicates that DeletedAt is a pointer to a string, meaning it can be nil if the device has not been deleted.
}

func (m *DeviceModel) ToDto() DeviceModelDto {
	var deviceModelDto DeviceModelDto

	// Custom copying logic for non-directly mappable fields
	copier.CopyWithOption(&deviceModelDto, &m, copier.Option{IgnoreEmpty: true})
	//  copies all values from m to deviceModelDto, ignoring any fields that are empty in m (i.e. zero values).
	deviceModelDto.CreatedAt = m.CreatedAt.Format("2006-01-02 15:04:05")
	deviceModelDto.UpdatedAt = m.UpdatedAt.Format("2006-01-02 15:04:05")
	if !m.DeletedAt.Time.IsZero() {
		d := m.DeletedAt.Time.Format("2006-01-02 15:04:05")
		deviceModelDto.DeletedAt = &d
	}

	return deviceModelDto
}
