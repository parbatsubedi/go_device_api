package repository

import (
	"go_api/database"
	"go_api/interfaces"
	"go_api/models"

	"gorm.io/gorm"
)

type DeviceCommandRepository struct {
	db *gorm.DB
}

var _ interfaces.IDeviceCommandRepository = (*DeviceCommandRepository)(nil)

func NewDeviceCommandRepository() *DeviceCommandRepository {
	return &DeviceCommandRepository{db: database.RootDatabase.DB}
}

// Create saves a new device command
func (r *DeviceCommandRepository) Create(command *models.DeviceCommandModel) error {
	return r.db.Create(command).Error
}

// FindByID finds a device command by ID
func (r *DeviceCommandRepository) FindByID(id uint) (models.DeviceCommandModel, bool) {
	var command models.DeviceCommandModel
	err := r.db.First(&command, id).Error
	if err != nil {
		return command, false
	}
	return command, true
}

// FindByDeviceID finds all commands for a specific device
func (r *DeviceCommandRepository) FindByDeviceID(deviceID uint) []models.DeviceCommandModel {
	var commands []models.DeviceCommandModel
	r.db.Where("device_id = ?", deviceID).Order("created_at DESC").Find(&commands)
	return commands
}

// FindPendingCommands finds all pending commands for a device
func (r *DeviceCommandRepository) FindPendingCommands(deviceID uint) []models.DeviceCommandModel {
	var commands []models.DeviceCommandModel
	r.db.Where("device_id = ? AND status = ?", deviceID, "pending").Find(&commands)
	return commands
}

// UpdateStatus updates the status of a command
func (r *DeviceCommandRepository) UpdateStatus(id uint, status string) error {
	return r.db.Model(&models.DeviceCommandModel{}).Where("id = ?", id).Update("status", status).Error
}

// MarkAsSent marks a command as sent
func (r *DeviceCommandRepository) MarkAsSent(id uint) error {
	return r.db.Model(&models.DeviceCommandModel{}).Where("id = ?", id).Update("status", "sent").Error
}

// MarkAsAcknowledged marks a command as acknowledged
func (r *DeviceCommandRepository) MarkAsAcknowledged(id uint) error {
	return r.db.Model(&models.DeviceCommandModel{}).Where("id = ?", id).Updates(map[string]interface{}{
		"status":   "acknowledged",
		"acked_at": gorm.Expr("NOW()"),
	}).Error
}
