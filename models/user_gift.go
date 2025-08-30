package models

import "github.com/google/uuid"

type UserGift struct {
	BaseModel
	GifterName   string     `json:"gifterName"`
	Message      string     `json:"message"`
	GiftID       *uuid.UUID `gorm:"type:uuid;index" json:"giftId"`
	Gift         Gift       `gorm:"foreignKey:GiftID;references:ID" json:"gift"`
	CustomAmount *float64   `json:"customAmount"`
}
