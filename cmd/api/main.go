package main

import (
	"Blogs_Backend/internal/database"
	"Blogs_Backend/internal/server"
	
)
func main(){
	database.ConnectDB()
	server.StartServer()
}