package controller

import (
	"Blogs_Backend/internal/database"
	"Blogs_Backend/internal/model"
	"Blogs_Backend/internal/utils"
	"encoding/json"
	"fmt"
	"net/http"
)

// Create User
func CreateUser(w http.ResponseWriter, r *http.Request) {
	var user model.User
	salt := utils.GenerateSalt()
	fmt.Println("Entered to function CreateUser")
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, `{"errorInvalid JSON input": "`+err.Error()+`"}`, http.StatusInternalServerError)
		return
	}
	user.Password = utils.HashPassword(user.Password, salt)
	user.Salt = salt

	database.DB.Create(&user)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "User created",
		"user":    user,
	})
}

// Login (Password verification) // email and password
func LoginUser(res http.ResponseWriter, req *http.Request) {
	var inputUser model.User
	var foundUser model.User

	if err := json.NewDecoder(req.Body).Decode(&inputUser); err != nil {
		http.Error(res, `{"error Invalid JSON input": "`+err.Error()+`"}`, http.StatusInternalServerError) // res,error string,num status
		return
	}
	if err := database.DB.Where("email=?", inputUser.Email).Find(&foundUser).Error; err != nil {
		http.Error(res, "user not found", http.StatusBadRequest)
		return
	}
	if utils.HashPassword(inputUser.Password, foundUser.Salt) != foundUser.Password {
		http.Error(res, "Invalied password", http.StatusUnauthorized)
		return
	}

	// sucessful message
	res.Header().Set("Content-Type", "applicationn/json")
	json.NewEncoder(res).Encode(map[string]interface{}{
		"message": "Login Sucessfull",
		"user":    foundUser,
	})
}
