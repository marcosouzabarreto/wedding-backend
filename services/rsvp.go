package services

import (
	"gorm.io/gorm"
	"wedding-backend/models"
)

type RSVPService struct {
	db *gorm.DB
}

func NewRSVPService(db *gorm.DB) *RSVPService {
	return &RSVPService{
		db: db,
	}
}

func (s *RSVPService) Create(rsvp *models.RSVP) (*models.RSVP, error) {
	if err := s.db.Create(&rsvp).Error; err != nil {
		return nil, err
	}

	return rsvp, nil
}
