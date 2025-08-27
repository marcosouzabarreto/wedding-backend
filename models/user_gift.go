package models

type UserGift struct {
	BaseModel
	GifterName string `json:"gifterName"`
	Message    string `json:"message"`
	GiftID     uint   `gorm:"not null;index" json:"giftId"`
	Gift       Gift   `gorm:"foreignKey:GiftID" json:"gift"`
}
