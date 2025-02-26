package database

import (
	"fmt"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

// User Model
type User struct {
	ID        uint       `gorm:"primaryKey"`
	Name      string
	Email     string     `gorm:"uniqueIndex"`
	Password  string
	CreatedAt time.Time
	
}

// Post Model
type Post struct {
	ID        uint       `gorm:"primaryKey"`
	Title     string
	Content   *string
	UserId    uint       `gorm:"not null"`
	CreatedAt time.Time
	Comments  []Comment  `gorm:"foreignKey:PostId"`
}

// Comment Model
type Comment struct {
	ID          uint      `gorm:"primaryKey"`
	PostId      uint      `gorm:"not null"`
	UserId      uint      `gorm:"not null;index"`
	CommentText string
	CreatedAt   time.Time
	Post        Post     `gorm:"foreignKey:PostId"`
	User        User     `gorm:"foreignKey:UserId"`
}

// ConnectDB initializes the PostgreSQL database connection
func ConnectDB() {
	dsn := "host=localhost user=postgres password=root dbname=Blogs port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to the database")
	}

	DB = db
	fmt.Println("Database connected successfully!")

	// Run Migrations
	db.AutoMigrate(&User{}, &Post{}, &Comment{})
	fmt.Println("Tables are created")
}
