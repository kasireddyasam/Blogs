package database

import (
	"Blogs_Backend/internal/entities"

	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

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
	db.AutoMigrate(&entities.User{}, &entities.Post{}, &entities.Comment{})
	fmt.Println("Tables are created")
}
