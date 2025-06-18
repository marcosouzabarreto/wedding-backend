package models

import "github.com/satori/go.uuid"

type Guest struct {
	BaseModel
	Name        string    `gorm:"not null" json:"name" validate:"required"`
	Phone       string    `json:"phone,omitempty"`
	FamilyID    uuid.UUID `gorm:"not null;index" json:"familyId"`
	Family      Family    `gorm:"foreignKey:FamilyID" json:"family"`
	RSVP        *RSVP     `gorm:"foreignKey:GuestID" json:"rsvp,omitempty"`
	IsMainGuest bool      `gorm:"default:false" json:"isMainGuest"`
}

type GuestInput struct {
	Name        string `json:"name" validate:"required"`
	Phone       string `json:"phone,omitempty"`
	FamilyID    string `json:"familyId" validate:"required"`
	IsMainGuest bool   `json:"isMainGuest,omitempty"`
}
