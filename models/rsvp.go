package models

import (
	"github.com/google/uuid"
	"time"
)

type RSVP struct {
	BaseModel
	WillAttend          bool      `json:"willAttend"`
	SpecialRequests     string    `json:"specialRequests,omitempty"`
	Message             string    `json:"message,omitempty"`
	DietaryRestrictions string    `json:"dietaryRestrictions,omitempty"`
	PlusOneCount        int       `gorm:"default:0" json:"plusOneCount"`
	ResponseDate        time.Time `gorm:"autoUpdateTime" json:"responseDate"`
	GuestID             uuid.UUID `gorm:"not null;unique;index" json:"guestId"`
	Guest               Guest     `gorm:"foreignKey:GuestID" json:"-"`
}

type RSVPInput struct {
	WillAttend          bool   `json:"willAttend"`
	SpecialRequests     string `json:"specialRequests,omitempty"`
	DietaryRestrictions string `json:"dietaryRestrictions,omitempty"`
	PlusOneCount        int    `json:"plusOneCount,omitempty"`
	GuestID             string `json:"guestId" validate:"required"`
	Message             string `json:"message,omitempty"`
}

type FamilyRSVPMemberInput struct {
	GuestID             string `json:"guestId" validate:"required"`
	WillAttend          bool   `json:"willAttend"`
	DietaryRestrictions string `json:"dietaryRestrictions,omitempty"`
}

type FamilyRSVPRequest struct {
	FamilyToken     string                  `json:"familyToken" validate:"required"`
	Email           string                  `json:"email"`
	Phone           string                  `json:"phone"`
	SpecialRequests string                  `json:"specialRequests,omitempty"`
	Message         string                  `json:"message,omitempty"`
	Guests          []FamilyRSVPMemberInput `json:"guests" validate:"required,min=1"`
}
