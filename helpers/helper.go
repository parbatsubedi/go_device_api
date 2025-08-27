package helpers

import "github.com/gin-gonic/gin"

func GetCurrentUserID(c *gin.Context) uint {
	userID, exists := c.Get("user_id")
	if !exists {
		return 0 // Return 0 if user ID is not found in context
	}
	
	// Convert the user ID to uint
	if id, ok := userID.(uint); ok {
		return id
	}
	
	// Handle case where user ID might be stored as float64 (from JWT claims)
	if idFloat, ok := userID.(float64); ok {
		return uint(idFloat)
	}
	
	return 0 // Return 0 if conversion fails
}
