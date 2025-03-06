package model

import (
	"Blogs_Backend/internal/database"
	"Blogs_Backend/internal/entities"
)

func (r *RepoImplement) CreateUser(user entities.User) error {
	err := database.DB.Create(&user).Error
	return err
}

func (r *RepoImplement) LoginUser(user entities.User) (entities.User, error) {
	err := database.DB.First(&user).Error
	return user, err
}
