package server

import (
	"Blogs_Backend/internal/entities"
)

func (p *ServiceImpl) CreateComment(comment entities.Comment) error {
	return p.Repo.CreateComment(comment)
}
