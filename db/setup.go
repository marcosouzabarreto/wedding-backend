package db

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"wedding-backend/models"
)

func InitDB() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	if err = db.AutoMigrate(
		&models.Family{},
		&models.Guest{},
		&models.RSVP{},
		&models.Gift{},
	); err != nil {
		return nil, err
	}

	return db, nil
}
