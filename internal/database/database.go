package database

import (
	"fmt"
	"log"
	"os"

	"github.com/BerylCAtieno/quickfire-backend/internal/models"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	var err error

	err = godotenv.Load()
	if err != nil {
		log.Printf("Warning: .env file not found or could not be loaded in ConnectDatabase: %v. Assuming environment variables are set externally.\n", err)
	}

	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")
	sslmode := os.Getenv("DB_SSLMODE")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=Africa/Nairobi",
		host, user, password, dbname, port, sslmode)

	// Opening PostgreSQL connection

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database: " + err.Error())
	}

	// Checking Connection
	sqlDB, err := DB.DB()
	if err != nil {
		panic("Failed to get generic database object from GORM: " + err.Error())
	}
	err = sqlDB.Ping()
	if err != nil {
		panic("Failed to ping database: " + err.Error())
	}

	// Migrate models
	err = DB.AutoMigrate(
		&models.User{},
		&models.Category{},
		&models.Quiz{},
		&models.Question{},
		&models.Answer{},
	)
	if err != nil {
		panic("Failed to auto migrate database: " + err.Error())
	}
	fmt.Println("Database connection established and migrations run successfully!")
}
