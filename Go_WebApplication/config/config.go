package config

import (
	"Go_WebApplication/logger"
	"Go_WebApplication/models"

	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	loggers "gorm.io/gorm/logger"
)

var DB *gorm.DB

func ConnectDatabase() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Retrieve variables
	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")
	sslmode := os.Getenv("DB_SSLMODE")

	// Construct DSN
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		host, user, password, dbname, port, sslmode,
	)

	// GORM logger
	newLogger := loggers.New(
		log.New(os.Stdout, "\r\n[GORM] ", log.LstdFlags), // Output to console
		loggers.Config{
			SlowThreshold:             time.Second,
			LogLevel:                  loggers.Info, // Show all SQL queries
			IgnoreRecordNotFoundError: true,
			Colorful:                  true,
		},
	)
	// Initialize database connection
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{Logger: newLogger})
	if err != nil {
		logger.Log.Error().Err(err).Msg("Failed to connect to the database:")
		log.Fatal("Failed to connect to the database:", err)
	}

	db.AutoMigrate(&models.User{}, &models.Patient{}) // Migrate User Table

	DB = db

}

// apisecret
func GetAPISecret() string {
	return os.Getenv("API_SECRET")
}
