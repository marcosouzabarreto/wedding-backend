package services

import (
	"gorm.io/gorm"
	"wedding-backend/models"
)

type UserGiftService struct {
	db *gorm.DB
}

func NewUserGiftService(db *gorm.DB) *UserGiftService {
	return &UserGiftService{db: db}
}

func (s *UserGiftService) GetAll() ([]models.UserGift, error) {
	var userGifts []models.UserGift
	if err := s.db.Preload("Gift").Find(&userGifts).Error; err != nil {
		return nil, err
	}
	return userGifts, nil
}

func (s *UserGiftService) Create(userGift models.UserGift) (models.UserGift, error) {
	if err := s.db.Create(&userGift).Error; err != nil {
		return models.UserGift{}, err
	}
	return userGift, nil
}
