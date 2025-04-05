package models

type RSVP struct {
	ID         string `gorm:"primaryKey;autoIncrement" json:"id"`
	WillAttend bool   `json:"willAttend"`
}
