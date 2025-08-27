package services

import (
	"errors"
	"github.com/google/uuid"
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

func (s *RSVPService) CreateFamilyRSVP(req *models.FamilyRSVPRequest) error {
	var family models.Family
	if err := s.db.Where("token = ?", req.FamilyToken).First(&family).Error; err != nil {
		return errors.New("family not found")
	}

	// Update family contact info
	family.Email = req.ContactInfo.Email
	family.Phone = req.ContactInfo.Phone
	if err := s.db.Save(&family).Error; err != nil {
		return err
	}

	for _, member := range req.Members {
		guestID, err := uuid.Parse(member.GuestID)
		if err != nil {
			return errors.New("invalid guest ID")
		}

		var rsvp models.RSVP
		// Check if RSVP already exists for this guest
		if err := s.db.Where("guest_id = ?", guestID).First(&rsvp).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				// Create new RSVP
				rsvp = models.RSVP{
					GuestID:             guestID,
					WillAttend:          member.WillAttend,
					SpecialRequests:     req.SpecialRequests,
					DietaryRestrictions: member.DietaryRestrictions,
					PlusOneCount:        0, // Assuming plus one count is handled separately or not applicable here
				}
				if err := s.db.Create(&rsvp).Error; err != nil {
					return err
				}
			} else {
				return err
			}
		} else {
			// Update existing RSVP
			rsvp.WillAttend = member.WillAttend
			rsvp.SpecialRequests = req.SpecialRequests
			rsvp.DietaryRestrictions = member.DietaryRestrictions
			if err := s.db.Save(&rsvp).Error; err != nil {
				return err			}
		}
	}

	return nil
}
