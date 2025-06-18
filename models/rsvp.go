package models

import (
	"time"
	"github.com/satori/go.uuid"
)

type RSVP struct {
	BaseModel
	WillAttend          bool      `json:"willAttend"`
	SpecialRequests     string    `json:"specialRequests,omitempty"`
	DietaryRestrictions string    `json:"dietaryRestrictions,omitempty"`
	PlusOneCount        int       `gorm:"default:0" json:"plusOneCount"`
	ResponseDate        time.Time `gorm:"autoUpdateTime" json:"responseDate"`
	GuestID             uuid.UUID `gorm:"not null;unique;index" json:"guestId"`
	Guest               Guest     `gorm:"foreignKey:GuestID" json:"guest"`
}

type RSVPInput struct {
	WillAttend          bool   `json:"willAttend"`
	SpecialRequests     string `json:"specialRequests,omitempty"`
	DietaryRestrictions string `json:"dietaryRestrictions,omitempty"`
	PlusOneCount        int    `json:"plusOneCount,omitempty"`
	GuestID             string `json:"guestId" validate:"required"`
}
