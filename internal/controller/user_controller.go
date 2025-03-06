package controller

import (
	"Blogs_Backend/internal/entities"
	"Blogs_Backend/internal/utils"
	"encoding/json"
	"net/http"
)

// type UserHandlerImpl struct {
// 	userService server.UserService
// }

// // Constructor function for UserHandler
// func NewUserHandler(service server.UserService) *UserHandlerImpl {
// 	return &UserHandlerImpl{
// 		userService: server.NewUserService(),
// 	}
// }

// Create User
func (h *PostHandlerImpl) CreateUser(w http.ResponseWriter, r *http.Request) {
	var user entities.User

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, `{"errorInvalid JSON input": "`+err.Error()+`"}`, http.StatusInternalServerError)
		return
	}

	if err := utils.Validator.Struct(user); err != nil {
		http.Error(w, "fields values can't be null", 500)
		return
	}

	if err := h.postService.CreateUser(user); err != nil {
		http.Error(w, "internal server issu", http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "User created",
		"user":    user,
	})
}

// Login (Password verification) // email and password
func (h *PostHandlerImpl) LoginUser(res http.ResponseWriter, req *http.Request) {
	var user entities.User

	if err := json.NewDecoder(req.Body).Decode(&user); err != nil {
		http.Error(res, `{"errorInvalid JSON input": "`+err.Error()+`"}`, http.StatusInternalServerError)
		return
	}

	if err := utils.Validator.Var(user.Email, "required,email"); err != nil {
		http.Error(res, "fields values can't be null", 500)
		return
	}
	user, err := h.postService.LoginUser(user)
	if err != nil {
		http.Error(res, "internal server issu", http.StatusInternalServerError)
		return 
	}

	// sucessful message
	res.Header().Set("Content-Type", "applicationn/json")
	json.NewEncoder(res).Encode(map[string]interface{}{
		"message": "Login Sucessfull",
		"user":    user,
	})
}
