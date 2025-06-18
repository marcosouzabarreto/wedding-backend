package services

import (
	"gorm.io/gorm"
	"wedding-backend/models"
)

type GuestService struct {
	db *gorm.DB
}

func NewGuestService(db *gorm.DB) *GuestService {
	return &GuestService{db: db}
}

func (s *GuestService) GetAll() ([]models.Guest, error) {
	var guests []models.Guest
	if err := s.db.Preload("Family").Preload("RSVP").Find(&guests).Error; err != nil {
		return nil, err
	}
	return guests, nil
}

func (s *GuestService) Create(guest models.Guest) (models.Guest, error) {
	if err := s.db.Create(&guest).Error; err != nil {
		return models.Guest{}, err
	}
	return guest, nil
}

func (s *GuestService) GetByID(id string) (models.Guest, error) {
	var guest models.Guest
	if err := s.db.Preload("Family").Preload("RSVP").First(&guest, id).Error; err != nil {
		return models.Guest{}, err
	}
	return guest, nil
}
