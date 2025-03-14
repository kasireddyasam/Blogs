package controller

import (
	"Blogs_Backend/internal/database"
	"Blogs_Backend/internal/entities"
	"Blogs_Backend/internal/server"
	"Blogs_Backend/internal/utils"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type PostHandlerImpl struct {
	postService server.PostService
}

func NewPostController(service server.PostService) *PostHandlerImpl {
	return &PostHandlerImpl{
		postService: server.NewService(),
	}
}

// Create Post
func (h *PostHandlerImpl) CreatePost(w http.ResponseWriter, r *http.Request) {
	var post entities.Post

	// Decode req body
	if err := json.NewDecoder(r.Body).Decode(&post); err != nil {
		http.Error(w, "Invalid JSON input", http.StatusBadRequest)
		return
	}

	// Validate user data
	if err := utils.Validator.Var(post.UserID, "required"); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// send object to service layer
	h.postService.CreatePost(post)

	// return response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "Post created",
		"post":    post,
	})
}

// Create Comment
func (h *PostHandlerImpl) CreateComment(res http.ResponseWriter, req *http.Request) {
	var comment entities.Comment
	// Decode req body
	if err := json.NewDecoder(req.Body).Decode(&comment); err != nil {
		http.Error(res, "Invalid JSON input", http.StatusBadRequest)
	}

	// Validate user data
	if err := utils.Validator.Struct(comment); err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}
	// send object to service layer
	if err := h.postService.CreateComment(comment); err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}

	// return response
	res.Header().Set("Content-Type", "application/json")
	json.NewEncoder(res).Encode(map[string]interface{}{
		"message": "Comment created",
		"comment": comment,
	})
}

// Get all Users
func (h *PostHandlerImpl) GetAllUsers(res http.ResponseWriter, req *http.Request) {
	var users []entities.User
	// decode req body
	// validation not required
	// database
	if database.DB == nil {
		http.Error(res, "Database not initilized", http.StatusInternalServerError)
	}

	// call to service layer
	users, err := h.postService.GetAllUsers()
	if err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}
	//.DB.Find(&users)
	res.Header().Set("Content-Type", "application/json")
	json.NewEncoder(res).Encode(users)
}

// find user by email
func (h *PostHandlerImpl) FindUserByEmail(res http.ResponseWriter, req *http.Request) {
	var user entities.User
	if err := json.NewDecoder(req.Body).Decode(&user); err != nil {
		http.Error(res, "invalied url", http.StatusBadRequest)
		return 
	}
	if err:=utils.Validator.Var(user.Email,"required");err!=nil{
		http.Error(res,"invalied email",http.StatusBadRequest)
		return
	}
	h.postService.FindUserByEmail(user.Email)
}

// Get All Posts
func (h *PostHandlerImpl) GetAllPosts(res http.ResponseWriter, req *http.Request) {
	//var posts []entities.Post
	// body check
	// validation not need
	// database initilization check
	if database.DB == nil {
		http.Error(res, "Database not initialized", 500)
		return
	}
	// call to server
	posts, err := h.postService.GetAllPosts()
	if err != nil {
		http.Error(res, err.Error(), 404)
	}
	database.DB.Find(&posts)
	res.Header().Set("Content-Type", "application/json")
	json.NewEncoder(res).Encode(posts)
}

// Get Post by ID
func (h *PostHandlerImpl) GetPostByID(res http.ResponseWriter, req *http.Request) {
	var post entities.Post
	// body check
	if err := json.NewDecoder(req.Body).Decode(&post); err != nil {
		http.Error(res, "invalid Json File", http.StatusBadRequest)
	}
	// validator
	if err := utils.Validator.Var(post.ID, "required"); err != nil {
		http.Error(res, "id is should not be nil", 400)
		return
	}

	id, error := strconv.Atoi(chi.URLParam(req, "id"))
	if error != nil {
		http.Error(res, "invalied type of id", 400)
	}
	NewId := uint(id)

	// call server
	post, err := h.postService.GetPostByID(NewId)
	if err != nil {
		http.Error(res, fmt.Sprintf("Not found with id %d", id), 404)
	}

	res.Header().Set("Content-Type", "application/json")
	json.NewEncoder(res).Encode(post)
}

// Get All Posts by a Specific User
func (h *PostHandlerImpl) GetUserPosts(res http.ResponseWriter, req *http.Request) {
	var post entities.Post
	var posts []entities.Post
	if err := json.NewDecoder(req.Body).Decode(&post); err != nil {
		http.Error(res, "invalied request ", 400)
		return
	}
	userID, error := strconv.Atoi(chi.URLParam(req, "user_id"))
	if error != nil {
		http.Error(res, fmt.Sprintf("Invalied UserId %d", userID), 400)
	}
	// validation
	if err := utils.Validator.Var(post.UserID, "required"); err != nil {
		http.Error(res, "input fields are missing", 400)
		return
	}
	postId, error := strconv.Atoi(chi.URLParam(req, "id"))
	if error != nil {
		http.Error(res, "invalied postId", 400)
		return
	}
	posts, err := h.postService.GetUserPosts(uint(postId))
	if err != nil {
		http.Error(res, "Server issu", 500)
		return
	}

	res.Header().Set("Content-Type", "application/json")
	json.NewEncoder(res).Encode(posts)
}

// Update Post
func (h *PostHandlerImpl) UpdatePost(res http.ResponseWriter, req *http.Request) {
	var post entities.Post
	if err := json.NewDecoder(req.Body).Decode(&post); err != nil {
		http.Error(res, "Invalid JSON input", http.StatusBadRequest)
		return
	}

	if err := utils.Validator.Struct(post); err != nil {
		http.Error(res, "input fields are missing", 400)
		return
	}

	err := h.postService.UpdatePost(post)
	if err != nil {
		http.Error(res, "Server issu", 500)
		return
	}

	res.Header().Set("Content-Type", "application/json")
	json.NewEncoder(res).Encode(map[string]interface{}{
		"message": "Post updated",
		"post":    post,
	})
}

// Delete Post
func (h *PostHandlerImpl) DeletePost(res http.ResponseWriter, req *http.Request) {
	var post entities.Post
	if err := json.NewDecoder(req.Body).Decode(&post); err != nil {
		http.Error(res, "Invalid JSON input", http.StatusBadRequest)
		return
	}

	if err := utils.Validator.Var(post.ID, "required"); err != nil {
		http.Error(res, "input fields are missing", 400)
		return
	}

	postId, error := strconv.Atoi(chi.URLParam(req, "id"))
	if error != nil {
		http.Error(res, "invalied postId", 400)
		return
	}
	err := h.postService.DeletePost(uint(postId))
	if err != nil {
		http.Error(res, "Server issu", 500)
		return
	}

	res.Header().Set("Content-Type", "application/json")
	json.NewEncoder(res).Encode(map[string]interface{}{
		"message":      "Post deleted",
		"Deleted Post": post,
	})
}
