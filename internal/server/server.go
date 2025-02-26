package server

import (
	"fmt"
	"net/http"
)

func StartServer() {
	r := SetupRoutes()
	fmt.Println("Server running on port 8080")
	http.ListenAndServe(":8080", r)
}
