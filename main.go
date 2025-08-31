package main

import (
	"log"
	"log/slog"
	"os"
	"path/filepath"
	"time"

	"go_api/database"
	"go_api/middlewares"
	"go_api/models"
	"go_api/repository"
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

	// Initialize logging
	setupLogging()

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
		serverHost = "localhost:8080" // Default value
	}

	publicDb := database.RootDatabase
	errMigratePublic := publicDb.DB.AutoMigrate(&models.UserModel{}, &models.DeviceModel{}, &models.DeviceLocationModel{}, &models.AuditTrailModel{}, &models.DeviceCommandModel{})
	if errMigratePublic != nil {
		panic("could not auto migrate")
		// return
	}

	// Initialize Routers
	router := gin.Default()                                   // Initialize the router
	router.Static("/static", "./static")                      // Serve static files
	router.StaticFile("/favicon.ico", "./static/favicon.ico") // Serve favicon
	router.LoadHTMLGlob(filepath.Join("templates", "**", "*"))

	// Auth Protected System Routes
	protectedSystemRoutes := router.Group("/api/")
	nonProtectedSystemRoutes := router.Group("/api/")
	adminRoutes := router.Group("/admin/")
	// Create user repository instance for authentication
	userRepo := repository.NewUserRepository()
	protectedSystemRoutes.Use(middlewares.RequireAuth(userRepo))

	// CORS Middleware
	router.Use(middlewares.CORSMiddleware())

	// Add Auth Routes
	routes.AddSystemAuthRoutes(protectedSystemRoutes)
	routes.AddNonAuthRoutes(nonProtectedSystemRoutes)
	routes.AddAdminRoutes(adminRoutes)

	// Start Server
	router.Run(":" + os.Getenv("PORT"))
}

func setupLogging() {
	logDir := "logs"
	currentDate := time.Now().Format("2006-01-02")
	logFile := filepath.Join(logDir, "app-"+currentDate+".log")

	// Create logs directory if it doesn't exist
	if _, err := os.Stat(logDir); os.IsNotExist(err) {
		os.Mkdir(logDir, 0755)
	}

	// Open log file for writing
	file, err := os.OpenFile(logFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err)
	}
	// Note: We don't close the file here to keep it open for logging

	handlerOpts := &slog.HandlerOptions{
		Level: slog.LevelDebug,
	}

	logger := slog.New(slog.NewTextHandler(file, handlerOpts))
	slog.SetDefault(logger)

	// Log rotation - create new files daily
	go func() {
		for {
			time.Sleep(24 * time.Hour) // Wait for a day
			createNewLogFile(logDir)
		}
	}()
}

func createNewLogFile(logDir string) {
	currentDate := time.Now().Format("2006-01-02")
	logFile := filepath.Join(logDir, "app-"+currentDate+".log")

	// Create new log file
	file, err := os.OpenFile(logFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Println("Error creating new log file:", err)
		return
	}
	defer file.Close()

	slog.Info("New daily log file created for " + currentDate)
}
