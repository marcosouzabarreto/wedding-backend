package models

type UserGift struct {
	BaseModel
	GifterName   string   `json:"gifterName"`
	Message      string   `json:"message"`
	GiftID       *uint    `gorm:"index" json:"giftId"`
	Gift         Gift     `gorm:"foreignKey:GiftID" json:"gift"`
	CustomAmount *float64 `json:"customAmount"`
}
