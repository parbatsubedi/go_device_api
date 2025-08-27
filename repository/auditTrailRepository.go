package repository

import (
	"go_api/database"
	"go_api/models"

	"gorm.io/gorm"
)

type AuditTrailRepository struct {
	db *gorm.DB
}

// NewAuditTrailRepository creates a new instance of AuditTrailRepository
func NewAuditTrailRepository() *AuditTrailRepository {
	return &AuditTrailRepository{
		db: database.RootDatabase.DB,
	}
}

func (r *AuditTrailRepository) Create(audit *models.AuditTrailModel) error {
	return r.db.Create(audit).Error
}

// FindByEntity finds audit trails for a specific entity
func (r *AuditTrailRepository) FindByEntity(entity string, entityID uint) ([]models.AuditTrailModel, error) {
	var audits []models.AuditTrailModel
	err := r.db.Where("entity = ? AND entity_id = ?", entity, entityID).
		Order("created_at DESC").
		Find(&audits).Error
	return audits, err
}

// FindByUser finds audit trails for a specific user
func (r *AuditTrailRepository) FindByUser(userID uint) ([]models.AuditTrailModel, error) {
	var audits []models.AuditTrailModel
	err := r.db.Where("user_id = ?", userID).
		Order("created_at DESC").
		Find(&audits).Error
	return audits, err
}

// FindByAction finds audit trails for a specific action type
func (r *AuditTrailRepository) FindByAction(action string) ([]models.AuditTrailModel, error) {
	var audits []models.AuditTrailModel
	err := r.db.Where("action = ?", action).
		Order("created_at DESC").
		Find(&audits).Error
	return audits, err
}

// FindRecent finds recent audit trails with limit
func (r *AuditTrailRepository) FindRecent(limit int) ([]models.AuditTrailModel, error) {
	var audits []models.AuditTrailModel
	err := r.db.Order("created_at DESC").
		Limit(limit).
		Find(&audits).Error
	return audits, err
}
