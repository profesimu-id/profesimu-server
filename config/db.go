package config

import (
	"fmt"
	"os"
	"profesimu/models"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Setup connection to database
func InitDatabase() *gorm.DB {
	err := godotenv.Load()
	if err != nil {
		panic("Failed to load env file")
	}

	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("postgres://%s:%s@%s:5432/%s", dbUser, dbPass, dbHost, dbName)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect database")
	}
	db.AutoMigrate(&models.User{}, &models.Profile{}, &models.Category{}, models.Profession{})

	return db
}

// Stop connection to database
func CloseDatabase(db *gorm.DB) {
	dbPostgres, err := db.DB()
	if err != nil {
		panic("Failed to close connection database")
	}
	dbPostgres.Close()
}
