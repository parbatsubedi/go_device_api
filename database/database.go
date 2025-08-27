package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *sql.DB

// func InitDB(cfg *config.Config) error {
// 	connectionString := fmt.Sprintf(
// 		"host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
// 		cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPassword, cfg.DBName, cfg.DBSSLMode,
// 	)

// 	var err error
// 	DB, err = sql.Open("postgres", connectionString)
// 	if err != nil {
// 		return fmt.Errorf("failed to open database connection: %v", err)
// 	}

// 	// Test the connection
// 	err = DB.Ping()
// 	if err != nil {
// 		return fmt.Errorf("failed to ping database: %v", err)
// 	}

//		log.Println("Successfully connected to PostgreSQL database")
//		return nil
//	}
type Database struct {
	Path string
	DB   *gorm.DB
}

var RootDatabase Database

func InitDB(path string) {
	var err error
	RootDatabase = Database{}
	RootDatabase.Path = path

	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		dbHost,
		dbPort,
		dbUser,
		dbPassword,
		dbName,
	)

	RootDatabase.DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	// RootDatabase.DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(fmt.Errorf("failed to connect database: %s", err))
	}

	fmt.Println("Database initialized", path)

}

func CloseDB() {
	if DB != nil {
		DB.Close()
		log.Println("Database connection closed")
	}
}
