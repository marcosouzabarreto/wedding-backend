package services

import (
	"wedding-backend/db"
	"wedding-backend/models"
)

type RSVPService struct{}

func NewRSVPService() *RSVPService {
	return &RSVPService{}
}

func (s *RSVPService) Create(rsvp *models.RSVP) (*models.RSVP, error) {
	if err := db.DB.Create(&rsvp).Error; err != nil {
		return nil, err
	}

	return rsvp, nil
}
