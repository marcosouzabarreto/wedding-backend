package models

type Family struct {
	BaseModel
	Name   string  `gorm:"unique;not null" json:"name" validate:"required"`
	Token  string  `gorm:"unique;not null" json:"token"`
	Email  string  `json:"email"`
	Phone  string  `json:"phone"`
	Guests []Guest `gorm:"foreignKey:FamilyID" json:"guests,omitempty"`
}

type FamilyInput struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
}
