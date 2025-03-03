package controller

import (
	"Blogs_Backend/internal/database"
	"Blogs_Backend/internal/model"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
)

// Create Post
func CreatePost(w http.ResponseWriter, r *http.Request) {
	var post model.Post
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
	var comment model.Comment

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

// Get all Users
func GetAllUsers(res http.ResponseWriter, req *http.Request) {
	if database.DB == nil {
		http.Error(res, "Database not initilalized", http.StatusInternalServerError)
	}
	var users []model.User
	database.DB.Find(&users)
	res.Header().Set("Content-Type", "application/json")
	json.NewEncoder(res).Encode(users)
}

// Get All Posts
func GetAllPosts(w http.ResponseWriter, r *http.Request) {
	if database.DB == nil {
		http.Error(w, "Database not initialized", http.StatusInternalServerError)
		return
	}
	var posts []model.Post
	database.DB.Find(&posts)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(posts)
}

// Get Post by ID
func GetPostByID(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	var post model.Post

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
	var posts []model.Post

	database.DB.Where("user_id = ?", userID).Find(&posts)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(posts)
}

// Update Post
func UpdatePost(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	var post model.Post

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
	var post model.Post

	if err := database.DB.Delete(&model.Post{}, id).Error; err != nil {
		http.Error(w, "Post not found", http.StatusNotFound)
		return
	}
	database.DB.First(&post, id)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message":      "Post deleted",
		"Deleted Post": post,
	})
}
