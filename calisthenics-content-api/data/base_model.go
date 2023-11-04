package data

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type BaseModel struct {
	ID         string         `gorm:"primaryKey;type:uuid;" json:"id"`
	CreatedAt  time.Time      `json:"created_date"`
	UpdatedAt  time.Time      `json:"last_modified_date"`
	RowVersion int64          `gorm:"default:0" json:"row_version"`
	DeletedAt  gorm.DeletedAt `gorm:"index" json:"-"`
}

func (base *BaseModel) BeforeCreate(_ *gorm.DB) (err error) {
	u := uuid.New().String()
	base.ID = u
	return
}

func (base *BaseModel) BeforeUpdate(_ *gorm.DB) (err error) {
	base.RowVersion++
	return
}
