package db

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"wedding-backend/models"
)

var DB *gorm.DB

func InitDB() error {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})

	if err != nil {
    return err
	}

	if err = db.AutoMigrate(&models.Guest{}, &models.Family{}, &models.RSVP{}); err != nil {
    return err
  }

	DB = db

	return nil
}
