package server

import (
	"Blogs_Backend/internal/entities"
	"Blogs_Backend/internal/model"
)

type ServiceImpl struct {
	Repo model.Repo
}

func (s *ServiceImpl) CreatePost(post entities.Post) error {
	//validation will keep in model
	return s.Repo.CreatePost(post)
}

func (s *ServiceImpl) GetAllPosts() ([]entities.Post, error) {
	return s.Repo.GetAllPosts()
}

func (s *ServiceImpl) GetPostByID(id uint) (entities.Post, error) {
	return s.Repo.GetPostByID(id)
}

func (s *ServiceImpl) GetUserPosts(userId uint) ([]entities.Post, error) {
	return s.Repo.GetUserPosts(userId)
}

func (s *ServiceImpl) UpdatePost(post entities.Post) error {
	return s.Repo.UpdatePost(post)
}

func (s *ServiceImpl) DeletePost(postId uint) error {
	return s.Repo.DeletePost(postId)
}

func (s *ServiceImpl) GetAllUsers() ([]entities.User, error) {
	return s.Repo.GetAllUsers()
}
func (s *ServiceImpl) FindUserByEmail(email string) (entities.User, error) {
	return s.Repo.FindUserByEmail(email)
}

// type PostOperations interface {
// 	CreatePost(post *entities.Post) error
// 	GenerateSalt() ([]entities.Post, error)
// 	GetPostByID(id uint) (*entities.Post, error)
// 	GetUserPosts(userId uint) ([]entities.Post, error)
// 	UpdatePost(post *entities.Post) error
// 	DeletePost(id uint) error
// 	GetAllPosts() ([]entities.Post,error)
// }

// type CommentOperations interface {
// 	CreateComment(comment *entities.Comment) error
// }
