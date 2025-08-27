package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	ServerPort    int
	ServerHost    string
	DBHost        string
	DBPort        int
	DBName        string
	DBUser        string
	DBPassword    string
	DBSSLMode     string
	AppSecret     string
	AppEnv        string
	SessionSecret string
	SessionMaxAge int
}

func LoadConfig() (*Config, error) {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: .env file not found, using system environment variables")
	}

	// Parse server configuration
	serverPort, err := strconv.Atoi(getEnv("SERVER_PORT", "8080"))
	if err != nil {
		return nil, fmt.Errorf("invalid SERVER_PORT: %v", err)
	}

	sessionMaxAge, err := strconv.Atoi(getEnv("SESSION_MAX_AGE", "3600"))
	if err != nil {
		return nil, fmt.Errorf("invalid SESSION_MAX_AGE: %v", err)
	}

	dbPort, err := strconv.Atoi(getEnv("DB_PORT", "5432"))
	if err != nil {
		return nil, fmt.Errorf("invalid DB_PORT: %v", err)
	}

	config := &Config{
		ServerPort:    serverPort,
		ServerHost:    getEnv("SERVER_HOST", "localhost"),
		DBHost:        getEnv("DB_HOST", "localhost"),
		DBPort:        dbPort, // Default PostgreSQL port
		DBName:        getEnv("DB_NAME", "postgres"),
		DBUser:        getEnv("DB_USER", "postgres"),
		DBPassword:    getEnv("DB_PASSWORD", "password"),
		DBSSLMode:     getEnv("DB_SSLMODE", "disable"),
		AppSecret:     getEnv("APP_SECRET", "your-secret-key-here"),
		AppEnv:        getEnv("APP_ENV", "development"),
		SessionSecret: getEnv("SESSION_SECRET", "session-secret-key"),
		SessionMaxAge: sessionMaxAge,
	}

	// Override DB_PORT if specified
	if portStr := os.Getenv("DB_PORT"); portStr != "" {
		dbPort, err := strconv.Atoi(portStr)
		if err != nil {
			return nil, fmt.Errorf("invalid DB_PORT: %v", err)
		}
		config.DBPort = dbPort
	}

	return config, nil
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
