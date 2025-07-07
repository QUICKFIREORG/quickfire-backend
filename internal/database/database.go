package database

import (
	"fmt"
	"log"
	"time"

	"github.com/BerylCAtieno/quickfire-backend/config"
	"github.com/BerylCAtieno/quickfire-backend/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		config.Config.DBHost,
		config.Config.DBUser,
		config.Config.DBPassword,
		config.Config.DBName,
		config.Config.DBPort,
		config.Config.DBSSLMode,
		config.Config.DBTimeZone,
	)

	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	sqlDB, err := DB.DB()
	if err != nil {
		log.Fatalf("Failed to get generic database object from GORM: %v", err)
	}

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(5 * time.Minute)

	err = sqlDB.Ping()
	if err != nil {
		log.Fatalf("Failed to ping database: %v", err)
	}

	err = DB.AutoMigrate(
		&models.User{},
		&models.Category{},
		&models.Quiz{},
		&models.Question{},
		&models.Answer{},
	)
	if err != nil {
		log.Fatalf("Failed to auto migrate database: %v", err)
	}
	fmt.Println("Database connection established and migrations run successfully!")
}
