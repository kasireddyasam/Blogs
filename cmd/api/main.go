package main

import (
	"Blogs_Backend/internal/database"
	"Blogs_Backend/internal/routers"
)

func main() {
	database.ConnectDB()

	routers.StartServer()
}
