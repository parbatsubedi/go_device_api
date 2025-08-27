package routes

import (
	"go_api/controllers"
	"go_api/controllers/auth"

	"github.com/gin-gonic/gin"
)

func AddNonAuthRoutes(nonProtectedSystemRoutes *gin.RouterGroup) {
	authController := auth.NewAuthController()
	deviceController := controllers.MakeDeviceController()

	nonProtectedSystemRoutes.POST("/login", authController.Login)

	nonProtectedSystemRoutes.POST("/devices", deviceController.Create)
	nonProtectedSystemRoutes.GET("/devices", deviceController.GetAll)
	nonProtectedSystemRoutes.GET("/devices/:id", deviceController.GetByID)
	nonProtectedSystemRoutes.POST("/devices/:id", deviceController.Update)
	nonProtectedSystemRoutes.DELETE("/devices/:id", deviceController.Delete)

}
