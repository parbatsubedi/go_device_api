package apiresponses

import (
	"go_api/models"
)

// DeviceCommandResponse represents the response for a device command
type DeviceCommandResponse struct {
	Data DeviceCommandDto `json:"data"`
}

// DeviceCommandListResponse represents the response for a list of device commands
type DeviceCommandListResponse struct {
	Data []DeviceCommandDto `json:"data"`
}

// DeviceCommandDto represents the DTO for device command
type DeviceCommandDto struct {
	ID          uint       `json:"id"`
	DeviceID    uint       `json:"device_id"`
	CommandType string     `json:"command_type"`
	CommandData string     `json:"command_data"`
	Status      string     `json:"status"`
	SentAt      *string    `json:"sent_at"`
	AckedAt     *string    `json:"acked_at"`
	CreatedAt   string     `json:"created_at"`
}

// NewDeviceCommandResponse creates a new device command response
func NewDeviceCommandResponse() DeviceCommandResponse {
	return DeviceCommandResponse{}
}

// ToDeviceCommandDto converts DeviceCommandModel to DeviceCommandDto
func ToDeviceCommandDto(m models.DeviceCommandModel) DeviceCommandDto {
	dto := DeviceCommandDto{
		ID:          m.ID,
		DeviceID:    m.DeviceID,
		CommandType: m.CommandType,
		CommandData: m.CommandData,
		Status:      m.Status,
		CreatedAt:   m.CreatedAt.Format("2006-01-02 15:04:05"),
	}

	if m.SentAt != nil {
		sentAt := m.SentAt.Format("2006-01-02 15:04:05")
		dto.SentAt = &sentAt
	}

	if m.AckedAt != nil {
		ackedAt := m.AckedAt.Format("2006-01-02 15:04:05")
		dto.AckedAt = &ackedAt
	}

	return dto
}
