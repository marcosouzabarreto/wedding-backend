package models

type Family struct {
	BaseModel
	Name   string  `gorm:"unique;not null" json:"name" validate:"required"`
	Token  string  `gorm:"unique;not null" json:"token"`
	Guests []Guest `gorm:"foreignKey:FamilyID" json:"guests,omitempty"`
}

type FamilyInput struct {
	Name string `json:"name"`
}
