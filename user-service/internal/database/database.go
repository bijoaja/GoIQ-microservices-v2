package database

import (
	"log"
	"path/filepath"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"github.com/bijoaja/GoIQ-microservices.v2/user-service/internal/models"
)

func Connect() *gorm.DB {
	dbPath := filepath.Join(".", "users.db")
	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect db: %v", err)
	}
	return db
}

func Migrate(db *gorm.DB) error {
	return db.AutoMigrate(&models.User{})
}
