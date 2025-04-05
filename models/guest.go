package models

type Guest struct {
	ID       string `gorm:"primaryKey;autoIncrement" json:"id"`
	Name     string `gorm:"unique" json:"name"`
	FamilyID uint   `gorm:"not null" json:"familyId"`
	Family   Family `gorm:"foreignKey:FamilyID" json:"family"`
	RSVPID   *uint   `json:"RSVPId"`
	RSVP     *RSVP   `gorm:"foreignKey:RSVPID" json:"RSVP"`
}
