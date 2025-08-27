package services

import (
	"go_api/models"
	"testing"
)

type MockAuditTrailRepository struct {
	audits []models.AuditTrailModel
}

func (m *MockAuditTrailRepository) Create(audit *models.AuditTrailModel) error {
	m.audits = append(m.audits, *audit)
	return nil
}

func TestLogCreate(t *testing.T) {
	repo := &MockAuditTrailRepository{}
	service := &AuditService{auditRepo: repo}

	err := service.LogCreate("devices", 1, nil, "127.0.0.1", "")
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if len(repo.audits) != 1 {
		t.Errorf("Expected 1 audit record, got %d", len(repo.audits))
	}
}

func TestLogUpdate(t *testing.T) {
	repo := &MockAuditTrailRepository{}
	service := &AuditService{auditRepo: repo}

	err := service.LogUpdate("devices", 1, nil, "127.0.0.1", "", "new value")
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if len(repo.audits) != 1 {
		t.Errorf("Expected 1 audit record, got %d", len(repo.audits))
	}
}

func TestLogDelete(t *testing.T) {
	repo := &MockAuditTrailRepository{}
	service := &AuditService{auditRepo: repo}

	err := service.LogDelete("devices", 1, nil, "127.0.0.1", "")
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if len(repo.audits) != 1 {
		t.Errorf("Expected 1 audit record, got %d", len(repo.audits))
	}
}
