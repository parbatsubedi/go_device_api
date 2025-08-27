package interfaces

import "go_api/models"

// IAuditTrailRepository defines the interface for audit trail data operations.
type IAuditTrailRepository interface {
	Create(audit *models.AuditTrailModel) error
}
