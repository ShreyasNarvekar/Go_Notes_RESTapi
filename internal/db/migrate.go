package db

import (
	"go-notes-service/internal/models"
	"log"

	"gorm.io/gorm"
)

func RunMigrations(database *gorm.DB) {
	if err := database.AutoMigrate(&models.Note{}, &models.Task{}); err != nil {
		log.Fatalf("could not run database migrations: %v", err)
	}
}
