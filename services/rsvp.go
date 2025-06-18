package services

import (
	"wedding-backend/models"
	"gorm.io/gorm"
)

type RSVPService struct{
	db *gorm.DB
}

func NewRSVPService() *RSVPService {
	return &RSVPService{}
}

func (s *RSVPService) Create(rsvp *models.RSVP) (*models.RSVP, error) {
	if err := s.db.Create(&rsvp).Error; err != nil {
		return nil, err
	}

	return rsvp, nil
}
