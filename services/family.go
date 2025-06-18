package services

import (
	"crypto/rand"
	"strings"
	"wedding-backend/models"
	"gorm.io/gorm"
)

type FamilyService struct {
	db *gorm.DB
}

func NewFamilyService(db *gorm.DB) *FamilyService {
	return &FamilyService{
		db: db,
	}
}

func (s *FamilyService) GetAll() ([]models.Family, error) {
	var families []models.Family
	if err := s.db.Find(&families).Error; err != nil {
		return nil, err
	}
	return families, nil
}

func (s *FamilyService) Create(name string) (models.Family, error) {
	token, err := s.generateUniqueToken(name)
	if err != nil {
		return models.Family{}, err
	}

	family := models.Family{
		Name:  name,
		Token: token,
	}

	if err := s.db.Create(&family).Error; err != nil {
		return models.Family{}, err
	}
	return family, nil
}

func (s *FamilyService) GetByID(id string) (models.Family, error) {
	var family models.Family
	if err := s.db.First(&family, id).Error; err != nil {
		return models.Family{}, err
	}
	return family, nil
}

func (s *FamilyService) generateUniqueToken(familyName string) (string, error) {
	const tokenLength = 4
	const charset = "0123456789"

	bytes := make([]byte, tokenLength)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}

	for i, b := range bytes {
		bytes[i] = charset[b%byte(len(charset))]
	}

	prefix := strings.ToUpper(familyName[:3])
	token := prefix + string(bytes)

	var existing models.Family
	if err := s.db.Where("token = ?", token).First(&existing).Error; err == nil {
		return s.generateUniqueToken(familyName)
	} else if err != gorm.ErrRecordNotFound {
		return "", err
	}

	return token, nil
}
