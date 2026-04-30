package database

import (
	"fmt"
	"log"

	"go-jwt-auth-boilerplate/config"
	"go-jwt-auth-boilerplate/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)
var DB *gorm.DB
func Connect(cfg *config.Config) {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		cfg.DBHost,
		cfg.DBUser,
		cfg.DBPassword,
		cfg.DBName,
		cfg.DBPort,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {	
		log.Fatal("Failed to connect to database:", err)
	}

     // Auto-migrate — creates users table automatically
	db.AutoMigrate(&models.User{})


     DB = db
	fmt.Println("Database connected and migrated ✅")
}