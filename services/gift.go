package services

import (
	"gorm.io/gorm"
	"wedding-backend/models"
)

type GiftService struct {
	db *gorm.DB
}

func NewGiftService(db *gorm.DB) *GiftService {
	return &GiftService{db: db}
}

func (s *GiftService) GetAll() ([]models.Gift, error) {
	var gifts []models.Gift
	if err := s.db.Find(&gifts).Error; err != nil {
		return nil, err
	}
	return gifts, nil
}

func (s *GiftService) Create(gift models.Gift) (models.Gift, error) {
	if err := s.db.Create(&gift).Error; err != nil {
		return models.Gift{}, err
	}
	return gift, nil
}

func (s *GiftService) GetByID(id string) (models.Gift, error) {
	var gift models.Gift
	if err := s.db.First(&gift, id).Error; err != nil {
		return models.Gift{}, err
	}
	return gift, nil
}

func (s *GiftService) Update(id string, updatedGift models.Gift) (models.Gift, error) {
	var gift models.Gift
	if err := s.db.First(&gift, id).Error; err != nil {
		return models.Gift{}, err
	}

	// Update the fields
	gift.Name = updatedGift.Name
	gift.Description = updatedGift.Description
	gift.GifterName = updatedGift.GifterName
	gift.Message = updatedGift.Message

	if err := s.db.Save(&gift).Error; err != nil {
		return models.Gift{}, err
	}
	return gift, nil
}

func (s *GiftService) Delete(id string) error {
	if err := s.db.Delete(&models.Gift{}, id).Error; err != nil {
		return err
	}
	return nil
}
