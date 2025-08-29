package services

import (
	"wedding-backend/models"

	"gorm.io/gorm"
)

type UserGiftService struct {
	db *gorm.DB
}

func NewUserGiftService(db *gorm.DB) *UserGiftService {
	return &UserGiftService{
		db: db,
	}
}

func (s *UserGiftService) Create(gifterName, message string, giftIDs []uint, customAmount float64) ([]models.UserGift, error) {
	var userGifts []models.UserGift
	for _, giftID := range giftIDs {
		id := giftID
		userGift := models.UserGift{
			GifterName: gifterName,
			Message:    message,
			GiftID:     &id,
		}
		userGifts = append(userGifts, userGift)
	}

	if customAmount > 0 {
		amount := customAmount
		userGift := models.UserGift{
			GifterName:   gifterName,
			Message:      message,
			CustomAmount: &amount,
		}
		userGifts = append(userGifts, userGift)
	}

	if len(userGifts) > 0 {
		if err := s.db.Create(&userGifts).Error; err != nil {
			return nil, err
		}
	}

	return userGifts, nil
}
