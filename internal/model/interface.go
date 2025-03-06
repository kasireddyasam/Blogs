package model

import (
	"Blogs_Backend/internal/entities"
)

type Repo interface {
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

// type CommentRepo interface {
// 	CreateComment(comment entities.Comment) error
// }

// type UserRepo interface {
// 	CreateUser(user entities.User) error
// 	LoginUser(user entities.User) (entities.User, error)
// }

// Constructor
func NewRepository() Repo {
	return &RepoImplement{}
}
