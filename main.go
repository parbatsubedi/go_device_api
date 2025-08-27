package main

import (
	"log"
	"log/slog"
	"os"
	"path/filepath"

	"go_api/database"
	"go_api/models"
	"go_api/routes"
	"go_api/seeders"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables
	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: .env file not found, using system environment variables")
	}

	// Initialize database
	database.InitDB("")
	// Register Seeders
	seeders.RegisterSeeders([]seeders.DatabaseSeeder{
		seeders.UserSeeder{},
	})
	// Run Seeders
	err = seeders.RunSeeders()
	if err != nil {
		slog.Error("Seeding Failed!", slog.Any("error", err))

	} else {
		slog.Info("Seeding successful!")
	}

	serverHost := os.Getenv("APP_URL") // Get host from environment variable
	if serverHost == "" {
		serverHost = "localhost:8881" // Default value
	}

	// SETUP Logging
	// You could set this to any `io.Writer` such as a file
	file, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	handlerOpts := &slog.HandlerOptions{
		Level: slog.LevelDebug,
	}

	logger := slog.New(slog.NewTextHandler(file, handlerOpts))
	slog.SetDefault(logger)
	slog.Debug("Main Application Starting Point")

	publicDb := database.RootDatabase
	errMigratePublic := publicDb.DB.AutoMigrate(&models.UserModel{}, &models.DeviceModel{})
	if errMigratePublic != nil {
		panic("could not auto migrate")
		// return
	}

	// Initialize Routers
	router := gin.Default()                                   // Initialize the router
	router.Static("/static", "./static")                      // Serve static files
	router.StaticFile("/favicon.ico", "./static/favicon.ico") // Serve favicon
	router.LoadHTMLGlob(filepath.Join("templates", "**", "*"))
	// router.GET("/", handlers.HandleDefaultPage())
	// router.Use(middlewares.CORSMiddleware())
	// router.Use(middlewares.ErrorHandler())
	// router.Use(middlewares.AttachMdcID)

	// Auth Protected System Routes
	protectedSystemRoutes := router.Group("/api/")
	nonProtectedSystemRoutes := router.Group("/api/")
	adminRoutes := router.Group("/admin/")
	// protectedSystemRoutes.Use(middlewares.RequireAuth)

	// Add Auth Routes
	routes.AddSystemAuthRoutes(protectedSystemRoutes)
	routes.AddNonAuthRoutes(nonProtectedSystemRoutes)
	routes.AddAdminRoutes(adminRoutes)

	// Start Server
	router.Run(":" + os.Getenv("PORT"))

}
