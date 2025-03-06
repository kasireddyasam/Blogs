package entities

import (
	"time"
)

// Comment Model
type Comment struct {
	ID        uint   `gorm:"primaryKey"`
	PostID    uint   `gorm:"not null"`
	Post      Post   `gorm:"foreignKey:PostID;constraint:OnDelete:CASCADE;"`
	UserID    uint   `gorm:"not null"`
	User      User   `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE;"`
	Content   string `validate:"required"`
	CreatedAt time.Time
}
