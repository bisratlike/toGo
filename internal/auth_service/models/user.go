package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
    ID        uuid.UUID `gorm:"type:uuid;primary_key;" json:"id"`
    FullName  string    `gorm:"uniqueIndex;not null" json:"full_name"`
    Email     string    `gorm:"uniqueIndex;not null" json:"email"`
    Password  string    `gorm:"not null" json:"-"`
    Role      string    `gorm:"default:user" json:"role"`

    CreatedAt time.Time  `json:"created_at"`
    UpdatedAt time.Time  `json:"updated_at"`
    DeletedAt *time.Time `gorm:"index" json:"-"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = uuid.New()
	return
}