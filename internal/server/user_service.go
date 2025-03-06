package server

import (
	"Blogs_Backend/internal/entities"
	"Blogs_Backend/internal/utils"
	"errors"
	"fmt"
)

func (p *ServiceImpl) CreateUser(user entities.User) error {
	user.Salt = utils.GenerateSalt()
	return p.Repo.CreateUser(user)
}

func (p *ServiceImpl) LoginUser(user entities.User) (entities.User, error) {
	email := user.Email
	password := user.Password
	FoundUser, err := p.Repo.FindUserByEmail(email)
	if err != nil {
		fmt.Printf("error from model")
		return entities.User{}, err // return empty user,err
	}

	if FoundUser.Password != utils.HashPassword(password, FoundUser.Salt) {
		return entities.User{}, errors.New("invalid password") // empty user
	}

	return FoundUser, nil //  Success case
}
