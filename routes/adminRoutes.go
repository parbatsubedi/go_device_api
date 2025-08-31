package routes

import (
	"net/http"

	"go_api/controllers"

	"github.com/gin-gonic/gin"
)

func AddAdminRoutes(adminRoutes *gin.RouterGroup) {
	deviceController := controllers.MakeDeviceController()

	// Root route that serves the dashboard with data
	adminRoutes.GET("/", func(c *gin.Context) {
		// Get dashboard data
		deviceController.GetDashboardData(c)
	})

	adminRoutes.GET("/dashboard", func(c *gin.Context) {
		// Get dashboard data
		dashboardData, err := deviceController.GetDashboardDataForTemplate(c)
		if err != nil {
			c.HTML(http.StatusOK, "views/dashboard.html", gin.H{
				"title": "Admin Dashboard",
				"active_devices": 0,
				"average_battery": 0,
				"alerts": 0,
				"offline_devices": 0,
				"devices": []interface{}{},
			})
			return
		}
		
		c.HTML(http.StatusOK, "views/dashboard.html", dashboardData)
	})

	// API endpoint for dashboard data
	adminRoutes.GET("/api/dashboard", deviceController.GetDashboardData)
}
