package routers

// initializing and configuring the router
//  sets up middleware, logging, CORS, and global settings.
// checking proteted router and public router
import (
	"fmt"
	"net/http"
)

func StartServer() {
	r := SetupRoutes()
	http.ListenAndServe(":8080", r)
	fmt.Println("Server running on port 8080")
}
