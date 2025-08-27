package services

import (
	"encoding/json"
	"go_api/models"
	"go_api/repository"
	"time"

	"github.com/gin-gonic/gin"
)

// Statically assert that *AuditService implements IAuditService.
var _ IAuditService = (*AuditService)(nil)

type AuditService struct {
	repo repository.IAuditTrailRepository
}

func NewAuditService(repo repository.IAuditTrailRepository) *AuditService {
	return &AuditService{
		repo: repo,
	}
}

// Log creates an audit trail entry.
func (s *AuditService) Log(c *gin.Context, entity string, entityID uint, action string, oldValue, newValue any) error {
	// Extract user info from context (set by auth middleware)
	userID, _ := c.Get("user_id")
	userIDUint, ok := userID.(uint)
	var userIDPtr *uint
	if ok {
		userIDPtr = &userIDUint
	}

	oldValueJSON, _ := json.Marshal(oldValue)
	newValueJSON, _ := json.Marshal(newValue)

	audit := &models.AuditTrailModel{
		UserID:    userIDPtr,
		Entity:    entity,
		EntityID:  entityID,
		Action:    action,
		OldValue:  string(oldValueJSON),
		NewValue:  string(newValueJSON),
		IPAddress: c.ClientIP(),
		Timestamp: time.Now(),
	}

	return s.repo.Create(audit)
}
