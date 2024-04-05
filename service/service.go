package service

import (
	"time"

	"github.com/google/uuid"
)

type BaseModel struct {
	ID        uuid.UUID  `json:"id" gorm:"type:uuid"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"-"`
	DeletedAt *time.Time `json:"-"`
}

type User struct {
	BaseModel
	Email     string `json:"email" binding:"required" gorm:"unique"`
	FirstName string `json:"first_name" `
	LastName  string `json:"last_name" `
	Password  string `json:"password" binding:"required"`
	IsToWed   bool   `json:"is_to_wed"`
}
