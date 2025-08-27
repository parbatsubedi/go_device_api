package routes

import (
	"go_api/controllers"

	"github.com/gin-gonic/gin"
)

func AddSystemAuthRoutes(protectedRoutes *gin.RouterGroup) {

	deviceController := controllers.MakeDeviceController()

	protectedRoutes.POST("/devices", deviceController.Create)
	protectedRoutes.GET("/devices", deviceController.GetAll)
	protectedRoutes.GET("/devices/:id", deviceController.GetByID)
	protectedRoutes.POST("/devices/:id", deviceController.Update)
	protectedRoutes.DELETE("/devices/:id", deviceController.Delete)

}
