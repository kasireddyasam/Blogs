package server

import (
	"Blogs_Backend/internal/entities"
	"Blogs_Backend/internal/model"
)

type PostService interface {
	GetAllUsers() ([]entities.User, error)
	FindUserByEmail(email string) (entities.User, error)
	CreatePost(post entities.Post) error
	GetPostByID(id uint) (entities.Post, error)
	GetUserPosts(userId uint) ([]entities.Post, error)
	UpdatePost(post entities.Post) error
	DeletePost(id uint) error
	GetAllPosts() ([]entities.Post, error)
	CreateUser(user entities.User) error
	LoginUser(user entities.User) (entities.User, error)
	CreateComment(comment entities.Comment) error
}

// Constructor
func NewService() PostService {
	return &ServiceImpl{
		Repo: model.NewRepository(),
	}
}
