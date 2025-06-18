package models

import (
	"github.com/satori/go.uuid"
	"gorm.io/gorm"
	"time"
)

type BaseModel struct {
	ID        uuid.UUID `gorm:"type:uuid;primary_key;"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
}

func (base *BaseModel) BeforeCreate(db *gorm.DB) error {
	uuid := uuid.NewV4()
	base.ID = uuid
	return nil
}
