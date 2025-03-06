package entities

import (
	"time"
)

// User Model
type User struct {
	ID        uint   `gorm:"primaryKey"`
	Name      string `validate:"required"`
	Email     string `validate:"required,email"`
	Password  string `validate:"required,min=5"`
	CreatedAt time.Time
	Salt      string
}
