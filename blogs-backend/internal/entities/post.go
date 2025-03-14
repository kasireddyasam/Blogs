package entities

import (
	"time"
)

type Post struct {
	ID        uint    `gorm:"primaryKey"`
	Title     string  `validate:"required"`
	Content   string `validate:"required"`
	UserID    uint    `gorm:"not null" validate:"required"`
	User      User    `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE;"`
	CreatedAt time.Time
	Comments  []Comment `gorm:"foreignKey:PostID;constraint:OnDelete:CASCADE;"`
}
