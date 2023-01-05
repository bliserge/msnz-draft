package model

import (
	"time"
)

type User struct {
	ID        string `gorm:"primary_key;"`
	Name      string `json:"name"`
	Phone     string `json:"phone" binding:"required"`
	Email     string `json:"email" binding:"required"`
	Password  string `json:"-"`
	Status    uint
	CreatedAt time.Time
	UpdatedAt time.Time
}
