package repository

import (
	"go_api/database"
	"go_api/models"

	"gorm.io/gorm"
)

type AuditTrailRepository struct {
	db *gorm.DB
}

// NewAuditTrailRepository initializes a new AuditTrailRepository
func NewAuditTrailRepository() *AuditTrailRepository {
	return &AuditTrailRepository{
		db: database.RootDatabase.DB,
	}
}

// Create saves a new audit trail record to the database
func (r *AuditTrailRepository) Create(audit *models.AuditTrailModel) error {
	result := r.db.Create(audit)
	return result.Error
}
