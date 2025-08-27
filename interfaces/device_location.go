package interfaces

import "go_api/models"

// IDeviceLocationRepository defines the interface for device location data operations.
type IDeviceLocationRepository interface {
	Create(deviceLocation *models.DeviceLocationModel) error
}
