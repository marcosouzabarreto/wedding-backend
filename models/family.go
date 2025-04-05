package models

type Family struct {
	ID    string `gorm:"primaryKey;autoIncrement" json:"id"`
	Name  string `gorm:"unique" json:"name"`
	Token string `gorm:"unique" json:"token"`
}

type FamilyInput struct {
	Name  string `json:"name"`
}
