package models

import (
	"log"
	"time"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type BaseModel struct {
	ID        uuid.UUID  `json:"id" gorm:"type:uuid;primary_key;"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt"`
	DeletedAt *time.Time `json:"deletedAt" sql:"index"`
}

func (base *BaseModel) BeforeCreate(db *gorm.DB) error {
	uuid, err := uuid.NewV7()
	if err != nil {
		log.Fatalf("Error generating random UUID")
	}

	base.ID = uuid
	return nil
}
