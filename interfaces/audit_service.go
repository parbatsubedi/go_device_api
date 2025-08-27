package interfaces

// IAuditService defines the interface for audit service operations.
type IAuditService interface {
	LogCreate(entity string, entityID uint, userID *uint, ipAddress string, newValue string) error
	LogUpdate(entity string, entityID uint, userID *uint, ipAddress string, oldValue string, newValue string) error
	LogDelete(entity string, entityID uint, userID *uint, ipAddress string, oldValue string) error
	LogAction(entity string, entityID uint, userID *uint, action string, ipAddress string, oldValue string, newValue string) error
}
