package model

import (
	"Blogs_Backend/internal/database"
	"Blogs_Backend/internal/entities"
)

// CreateComment inserts a new comment into the database
func (r *RepoImplement) CreateComment(comment entities.Comment) error {
	return database.DB.Create(comment).Error
}
