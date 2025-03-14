package model

import (
	"Blogs_Backend/internal/database"
	"Blogs_Backend/internal/entities"
)

type RepoImplement struct{}

// find user
func (r *RepoImplement) FindUserByEmail(email string) (entities.User, error) {
	var user entities.User
	err := database.DB.Where("email=?", email).First(&user).Error
	return user, err
}

// get all users
func (r *RepoImplement) GetAllUsers() ([]entities.User, error) {
	var users []entities.User
	err := database.DB.Find(&users).Error
	return users, err
}

// create Post
func (r *RepoImplement) CreatePost(post entities.Post) error {
	return database.DB.Create(&post).Error
}

// GetAllPosts GetPostByID
func (r *RepoImplement) GetAllPosts() ([]entities.Post, error) {
	var posts []entities.Post
	err := database.DB.Find(&posts).Error
	return posts, err
}

// GetAllPostsById
func (r *RepoImplement) GetPostByID(id uint) (entities.Post, error) {
	var post entities.Post
	err := database.DB.First(&post, id).Error
	return post, err
}

// GetUserPosts
func (r *RepoImplement) GetUserPosts(userID uint) ([]entities.Post, error) {
	var posts []entities.Post
	err := database.DB.Where("user_id = ?", userID).Find(&posts).Error
	return posts, err
}

// UpdatePost
func (r *RepoImplement) UpdatePost(post entities.Post) error {
	return database.DB.Save(post).Error
}

// DeletePost
func (r *RepoImplement) DeletePost(id uint) error {
	return database.DB.Delete(&entities.Post{}, id).Error
}
