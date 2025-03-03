package model

import (
	"time"
)

// User Model
type User struct {
	ID        uint `gorm:"primaryKey"`
	Name      string
	Email     string `gorm:"uniqueIndex"`
	Password  string
	CreatedAt time.Time
	Salt      string
}

// Post Model
type Post struct {
	ID        uint `gorm:"primaryKey"`
	Title     string
	Content   *string
	UserId    uint `gorm:"not null"`
	CreatedAt time.Time
	Comments  []Comment `gorm:"foreignKey:PostId"`
}

// Comment Model
type Comment struct {
	ID          uint `gorm:"primaryKey"`
	PostId      uint `gorm:"not null"`
	UserId      uint `gorm:"not null;index"`
	CommentText string
	CreatedAt   time.Time
	Post        Post `gorm:"foreignKey:PostId"`
	User        User `gorm:"foreignKey:UserId"`
}
