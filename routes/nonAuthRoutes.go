package routes

import (
	"go_api/controllers/auth"

	"github.com/gin-gonic/gin"
)

func AddNonAuthRoutes(nonProtectedSystemRoutes *gin.RouterGroup) {
	authController := auth.NewAuthController()

	nonProtectedSystemRoutes.POST("/login", authController.Login)

}
