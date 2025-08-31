package interfaces

import (
	"go_api/models"

	"github.com/gin-gonic/gin"
)

// IDeviceCommandRepository defines the interface for device command data operations.
type IDeviceCommandRepository interface {
	Create(command *models.DeviceCommandModel) error
	FindByID(id uint) (models.DeviceCommandModel, bool)
	FindByDeviceID(deviceID uint) []models.DeviceCommandModel
	FindPendingCommands(deviceID uint) []models.DeviceCommandModel
	UpdateStatus(id uint, status string) error
	MarkAsSent(id uint) error
	MarkAsAcknowledged(id uint) error
}

// IDeviceCommandController defines the interface for device command controller operations.
type IDeviceCommandController interface {
	SendCommand(c *gin.Context)
	GetDeviceCommands(c *gin.Context)
	GetPendingCommands(c *gin.Context)
	AcknowledgeCommand(c *gin.Context)
}
