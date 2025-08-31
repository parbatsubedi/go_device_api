package repository

import (
	"go_api/database"
	"go_api/models"

	"gorm.io/gorm"
)

type DeviceLocationRepository struct {
	db *gorm.DB
}

func NewDeviceLocationRepository() *DeviceLocationRepository {
	return &DeviceLocationRepository{
		db: database.RootDatabase.DB,
	}
}

func (r *DeviceLocationRepository) Create(deviceLocation *models.DeviceLocationModel) error {
	return r.db.Create(deviceLocation).Error
}

func (r *DeviceLocationRepository) FindLatestByDeviceID(deviceID uint) (models.DeviceLocationModel, error) {
	var location models.DeviceLocationModel
	err := r.db.Where("device_id = ?", deviceID).
		Order("recorded_at DESC").
		First(&location).Error
	
	// Return empty model instead of error for "record not found"
	if err != nil && err.Error() == "record not found" {
		return models.DeviceLocationModel{}, nil
	}
	
	return location, err
}
