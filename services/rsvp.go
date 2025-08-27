package services

import (
	"errors"
	"wedding-backend/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
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

func (s *RSVPService) CreateFamilyRSVP(req *models.FamilyRSVPRequest) error {
	var family models.Family
	if err := s.db.Where("token = ?", req.FamilyToken).First(&family).Error; err != nil {
		return errors.New("family not found")
	}

	family.Email = req.Email
	family.Phone = req.Phone

	if err := s.db.Save(&family).Error; err != nil {
		return err
	}

	for _, member := range req.Guests {
		guestID, err := uuid.Parse(member.GuestID)
		if err != nil {
			return errors.New("invalid guest ID")
		}

		var rsvp models.RSVP
		if err := s.db.Where("guest_id = ?", guestID).First(&rsvp).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				rsvp = models.RSVP{
					GuestID:             guestID,
					WillAttend:          member.WillAttend,
					SpecialRequests:     req.SpecialRequests,
					DietaryRestrictions: member.DietaryRestrictions,
					Message:             req.Message,
					PlusOneCount:        0,
				}
				if err := s.db.Create(&rsvp).Error; err != nil {
					return err
				}
			} else {
				return err
			}
		} else {
			rsvp.WillAttend = member.WillAttend
			rsvp.SpecialRequests = req.SpecialRequests
			rsvp.DietaryRestrictions = member.DietaryRestrictions
			rsvp.Message = req.Message
			if err := s.db.Save(&rsvp).Error; err != nil {
				return err
			}
		}
	}

	return nil
}
