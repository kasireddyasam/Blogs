package routers

import (
	"Blogs_Backend/internal/controller"
	"fmt"

	"github.com/go-chi/chi/v5"
)

// Setup API Routes using Chi
func SetupRoutes() *chi.Mux {
	r := chi.NewRouter()
	fmt.Println("New router is created")

	// User Routes
	r.Route("/user", func(r chi.Router) {
		r.Post("/signup", controller.CreateUser)
		r.Post("/login", controller.LoginUser)
		r.Get("/{user_id}/posts", controller.GetUserPosts)

	})

	// Post Routes
	r.Route("/posts", func(r chi.Router) {
		r.Post("/", controller.CreatePost)
		r.Get("/", controller.GetAllPosts)
		r.Get("/{id}", controller.GetPostByID)
		r.Put("/{id}", controller.UpdatePost)
		r.Delete("/{id}", controller.DeletePost)
	})

	// comments Routes
	r.Route("/comment", func(r chi.Router) {
		r.Post("/", controller.CreateComment)
	})

	r.Get("/users", controller.GetAllUsers)

	return r
}
