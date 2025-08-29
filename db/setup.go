package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"os"
	"wedding-backend/models"
)

func InitDB() (*gorm.DB, error) {
	var db *gorm.DB
	var err error

	databaseURL := os.Getenv("DATABASE_URL")
	if databaseURL != "" {
		db, err = gorm.Open(postgres.Open(databaseURL), &gorm.Config{})
	} else {
		db, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	}

	if err != nil {
		return nil, err
	}

	if err = db.AutoMigrate(
		&models.User{},
		&models.Family{},
		&models.Guest{},
		&models.RSVP{},
		&models.Gift{},
		&models.UserGift{},
	); err != nil {
		return nil, err
	}

	return db, nil
}
