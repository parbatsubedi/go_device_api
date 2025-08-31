package routes

import (
	"go_api/controllers"

	"github.com/gin-gonic/gin"
)

func AddSystemAuthRoutes(protectedRoutes *gin.RouterGroup) {

	deviceController := controllers.MakeDeviceController()
	deviceCommandController := controllers.MakeDeviceCommandController()

	protectedRoutes.POST("/devices", deviceController.Create)
	protectedRoutes.GET("/devices", deviceController.GetAll)
	protectedRoutes.GET("/devices/:id", deviceController.GetByID)
	protectedRoutes.POST("/devices/:id", deviceController.Update)
	protectedRoutes.DELETE("/devices/:id", deviceController.Delete)

	// Device Command Routes
	protectedRoutes.POST("/device-commands/send/:device_id", deviceCommandController.SendCommand)
	protectedRoutes.GET("/device-commands/device/:device_id", deviceCommandController.GetDeviceCommands)
	protectedRoutes.GET("/device-commands/pending/:device_id", deviceCommandController.GetPendingCommands)
	protectedRoutes.POST("/device-commands/acknowledge/:commandId", deviceCommandController.AcknowledgeCommand)

}
