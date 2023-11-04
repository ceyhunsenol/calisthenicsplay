package data

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	_ "gorm.io/gorm"
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

type User struct {
	BaseModel
	Name     string  `json:"name"`
	Username string  `json:"username"`
	Password string  `json:"password"`
	Email    string  `json:"email"`
	Profile  Profile `gorm:"foreignKey:UserID"`
}

type Profile struct {
	BaseModel
	UserID      string    `gorm:"foreignKey:UserID;" json:"user_id"`
	DateOfBirth time.Time `json:"date_of_birth"`
	AvatarURL   string    `json:"avatar_url"`
	Bio         string    `json:"bio"`
}
