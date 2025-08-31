package apirequests

// DeviceCommandRequest represents a request to send a command to a device
type DeviceCommandRequest struct {
	CommandType string `json:"command_type" form:"command_type" binding:"required"`
	CommandData string `json:"command_data" form:"command_data"`
}

// Supported command types
const (
	CommandTypeLock   = "lock"
	CommandTypeUnlock = "unlock"
	CommandTypeWipe   = "wipe"
	CommandTypeRing   = "ring"
	CommandTypeLocate = "locate"
	CommandTypeStatus = "status"
)
