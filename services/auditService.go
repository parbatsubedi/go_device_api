package services

import (
	"encoding/json"
	"go_api/interfaces"
	"go_api/models"
	"go_api/repository"
	"time"
)

type AuditService struct {
	auditRepo interfaces.IAuditTrailRepository
}

// NewAuditService creates a new instance of AuditService
func NewAuditService() *AuditService {
	return &AuditService{
		auditRepo: repository.NewAuditTrailRepository(),
	}
}

// LogCreate logs a create action to the audit trail
func (s *AuditService) LogCreate(entity string, entityID uint, userID *uint, ipAddress string, newValue string) error {
	return s.LogAction(entity, entityID, userID, "create", ipAddress, "", newValue)
}

// LogUpdate logs an update action to the audit trail
func (s *AuditService) LogUpdate(entity string, entityID uint, userID *uint, ipAddress string, oldValue string, newValue string) error {
	return s.LogAction(entity, entityID, userID, "update", ipAddress, oldValue, newValue)
}

// LogDelete logs a delete action to the audit trail
func (s *AuditService) LogDelete(entity string, entityID uint, userID *uint, ipAddress string, oldValue string) error {
	return s.LogAction(entity, entityID, userID, "delete", ipAddress, oldValue, "")
}

// LogAction logs a generic action to the audit trail
func (s *AuditService) LogAction(entity string, entityID uint, userID *uint, action string, ipAddress string, oldValue string, newValue string) error {
	audit := &models.AuditTrailModel{
		UserID:    userID,
		Entity:    entity,
		EntityID:  entityID,
		Action:    action,
		OldValue:  oldValue,
		NewValue:  newValue,
		IPAddress: ipAddress,
		Timestamp: time.Now(),
	}

	return s.auditRepo.Create(audit)
}

// LogCreateWithObject logs a create action with object serialization
func (s *AuditService) LogCreateWithObject(entity string, entityID uint, userID *uint, ipAddress string, newObj interface{}) error {
	newValue, err := json.Marshal(newObj)
	if err != nil {
		return err
	}
	return s.LogCreate(entity, entityID, userID, ipAddress, string(newValue))
}

// LogUpdateWithObjects logs an update action with object serialization
func (s *AuditService) LogUpdateWithObjects(entity string, entityID uint, userID *uint, ipAddress string, oldObj interface{}, newObj interface{}) error {
	oldValue, err := json.Marshal(oldObj)
	if err != nil {
		return err
	}
	newValue, err := json.Marshal(newObj)
	if err != nil {
		return err
	}
	return s.LogUpdate(entity, entityID, userID, ipAddress, string(oldValue), string(newValue))
}

// LogDeleteWithObject logs a delete action with object serialization
func (s *AuditService) LogDeleteWithObject(entity string, entityID uint, userID *uint, ipAddress string, oldObj interface{}) error {
	oldValue, err := json.Marshal(oldObj)
	if err != nil {
		return err
	}
	return s.LogDelete(entity, entityID, userID, ipAddress, string(oldValue))
}
