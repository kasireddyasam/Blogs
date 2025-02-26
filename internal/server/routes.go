package server

import (
	"Blogs_Backend/internal/database"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
)

// Create User
func CreateUser(w http.ResponseWriter, r *http.Request) {
	var user database.User

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Invalid JSON input", http.StatusBadRequest)
		return
	}

	database.DB.Create(&user)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "User created",
		"user":    user,
	})
}

// Create Post
func CreatePost(w http.ResponseWriter, r *http.Request) {
	var post database.Post

	if err := json.NewDecoder(r.Body).Decode(&post); err != nil {
		http.Error(w, "Invalid JSON input", http.StatusBadRequest)
		return
	}

	database.DB.Create(&post)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "Post created",
		"post":    post,
	})
}

// Create Comment
func CreateComment(w http.ResponseWriter, r *http.Request) {
	var comment database.Comment

	if err := json.NewDecoder(r.Body).Decode(&comment); err != nil {
		http.Error(w, "Invalid JSON input", http.StatusBadRequest)
		return
	}

	database.DB.Create(&comment)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "Comment created",
		"comment": comment,
	})
}

// Get All Posts
func GetAllPosts(w http.ResponseWriter, r *http.Request) {
	var posts []database.Post
	database.DB.Find(&posts)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(posts)
}

// Get Post by ID
func GetPostByID(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	var post database.Post

	if err := database.DB.First(&post, id).Error; err != nil {
		http.Error(w, "Post not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(post)
}

// Get All Posts by a Specific User
func GetUserPosts(w http.ResponseWriter, r *http.Request) {
	userID := chi.URLParam(r, "user_id")
	var posts []database.Post

	database.DB.Where("user_id = ?", userID).Find(&posts)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(posts)
}

// Update Post
func UpdatePost(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	var post database.Post

	if err := database.DB.First(&post, id).Error; err != nil {
		http.Error(w, "Post not found", http.StatusNotFound)
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&post); err != nil {
		http.Error(w, "Invalid JSON input", http.StatusBadRequest)
		return
	}

	database.DB.Save(&post)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "Post updated",
		"post":    post,
	})
}

// Delete Post
func DeletePost(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	if err := database.DB.Delete(&database.Post{}, id).Error; err != nil {
		http.Error(w, "Post not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "Post deleted",
	})
}

// Setup API Routes using Chi
func SetupRoutes() *chi.Mux {
	r := chi.NewRouter()

	r.Post("/users", CreateUser)
	r.Post("/posts", CreatePost)
	r.Post("/comments", CreateComment)
	r.Get("/posts", GetAllPosts)
	r.Get("/posts/{id}", GetPostByID)
	r.Get("/user/{user_id}/posts", GetUserPosts)
	r.Put("/posts/{id}", UpdatePost)
	r.Delete("/posts/{id}", DeletePost)

	return r
}
