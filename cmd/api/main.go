package main

import (
	"Blogs_Backend/internal/database"
	"Blogs_Backend/internal/routers"
	"log"
)

func main() {
	database.ConnectDB()
	if database.DB == nil {
		log.Fatal("Database connection failed")
	}
	routers.StartServer()
}
