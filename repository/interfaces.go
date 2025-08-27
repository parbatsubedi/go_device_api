package repository

import (
	"go_api/models"
)

// IUserRepository defines the interface for user data operations.
type IUserRepository interface {
	Save(m *models.UserModel) error
	FindAll() ([]models.UserModel, error)
	Exists(m models.UserModel) (bool, error)
	FindByID(id uint) (models.UserModel, error)
	FindByEmail(email string) (models.UserModel, error)
	FindByMobileNo(mobileNo string) (models.UserModel, error)
	PartialUpdate(m models.UserModel) error
	Delete(m models.UserModel) error
}

// IDeviceRepository defines the interface for device data operations.
type IDeviceRepository interface {
	Save(m *models.DeviceModel) error
	FindAll() ([]models.DeviceModel, error)
	FindByID(id uint) (models.DeviceModel, error)
	Exists(m models.DeviceModel) (bool, error)
	PartialUpdate(m *models.DeviceModel) error
	Delete(m models.DeviceModel) error
	FindByIMEI(imei string) (models.DeviceModel, error)
	FindByUserID(userID uint) ([]models.DeviceModel, error)
	FindByDeviceName(deviceName string) (models.DeviceModel, error)
}

// IDeviceLocationRepository defines the interface for device location data operations.
type IDeviceLocationRepository interface {
	Create(deviceLocation *models.DeviceLocationModel) error
}

type IAuditTrailRepository interface {
	Create(audit *models.AuditTrailModel) error
}
