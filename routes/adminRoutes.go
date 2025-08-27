package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func AddAdminRoutes(adminRoutes *gin.RouterGroup) {

	// Root route that serves the dashboard
	adminRoutes.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "views/dashboard.html", gin.H{
			"title": "Admin Dashboard",
		})
	})

	adminRoutes.GET("/dashboard", func(c *gin.Context) {
		c.HTML(http.StatusOK, "views/dashboard.html", gin.H{
			"title": "Admin Dashboard",
		})
	})
}
