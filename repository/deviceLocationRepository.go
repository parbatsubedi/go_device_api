package repository

import (
	"go_api/models"

	"gorm.io/gorm"
)

type DeviceLocationRepository struct {
	db *gorm.DB
}

func NewDeviceLocationRepository(db *gorm.DB) *DeviceLocationRepository {
	return &DeviceLocationRepository{db}
}

func (r *DeviceLocationRepository) Create(deviceLocation *models.DeviceLocationModel) error {
	return r.db.Create(deviceLocation).Error
}
