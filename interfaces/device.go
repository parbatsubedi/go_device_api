package interfaces

import "go_api/models"

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
