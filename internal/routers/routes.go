package routers

import (
	"Blogs_Backend/internal/controller"
	"Blogs_Backend/internal/server"
	"fmt"

	"github.com/go-chi/chi/v5"
)

// Setup API Routes using Chi
func SetupRoutes() *chi.Mux {
	r := chi.NewRouter()
	fmt.Println("New router is created")

	// Initialize Service Layer
	postService := server.NewService()

	// Initialize Controller with Service Layer
	Controller := controller.NewPostController(postService)

	// User Routes
	r.Route("/user", func(r chi.Router) {
		r.Post("/signup", Controller.CreateUser)
		r.Post("/login", Controller.LoginUser)

	})

	// Post Routes
	r.Route("/posts", func(r chi.Router) {
		r.Post("/", Controller.CreatePost)
		r.Get("/", Controller.GetAllPosts)
		r.Get("/{id}", Controller.GetPostByID)
		r.Put("/{id}", Controller.UpdatePost)
		r.Delete("/{id}", Controller.DeletePost)
	})

	// comments Routes
	r.Route("/comment", func(r chi.Router) {
		r.Post("/", Controller.CreateComment)
	})

	r.Get("/users", Controller.GetAllUsers)
	r.Get("/{user_id}/posts", Controller.GetUserPosts)

	return r
}
