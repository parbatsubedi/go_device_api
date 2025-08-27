package routes

import (
	"github.com/gin-gonic/gin"
)

func AddSystemAuthRoutes(protectedRoutes *gin.RouterGroup) {

	// userController := controllers.MakeUserController()
	// serviceTypesController := controllers.MakeServiceTypesController()
	// serviceController := controllers.NewServiceController()
	// bookingController := controllers.NewBookingController()

	// protectedRoutes.POST("/users", userController.Create)
	// protectedRoutes.GET("/users", userController.GetAll)
	// protectedRoutes.GET("/users/:id", userController.GetById)
	// protectedRoutes.PUT("/users/:id", userController.Update)
	// protectedRoutes.DELETE("/users/:id", userController.Delete)

	// protectedRoutes.POST("/service-type", serviceTypesController.Create)
	// protectedRoutes.GET("/service-type", serviceTypesController.GetAll)
	// protectedRoutes.GET("/service-type/:id", serviceTypesController.GetByID)
	// protectedRoutes.PUT("/service-type/:id", serviceTypesController.Update)
	// protectedRoutes.DELETE("/service-type/:id", serviceTypesController.Delete)

	// protectedRoutes.POST("/service", serviceController.Create)
	// protectedRoutes.GET("/service", serviceController.GetAll)
	// protectedRoutes.GET("/service/:id", serviceController.GetByID)
	// protectedRoutes.PUT("/service/:id", serviceController.Update)
	// protectedRoutes.DELETE("/service/:id", serviceController.Delete)

	// protectedRoutes.POST("/booking", bookingController.Create)
	// protectedRoutes.GET("/booking", bookingController.GetAll)
	// protectedRoutes.GET("/booking/:id", bookingController.GetByID)
	// protectedRoutes.PUT("/booking/:id", bookingController.Update)
	// protectedRoutes.DELETE("/booking/:id", bookingController.Delete)

}
