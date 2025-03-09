package server

import (
	"Blogs_Backend/internal/entities"
	"Blogs_Backend/internal/utils"
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)
// create user
func (p *ServiceImpl) CreateUser(user entities.User) error {
	user.Salt = utils.GenerateSalt()
	return p.Repo.CreateUser(user)
}

var jwtSecret = []byte("your_secret_key") // Use env variables in production!

// Function to generate JWT token
func generateJWT(user entities.User) (string, error) {
	claims := jwt.MapClaims{
		"user_id": user.ID,
		"email":   user.Email,
		"exp":     time.Now().Add(time.Hour * 24).Unix(), // Token expires in 24 hours
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

// Login function
func (p *ServiceImpl) LoginUser(user entities.User) (string, error) {
	email := user.Email
	password := user.Password

	// Find user in DB
	FoundUser, err := p.Repo.FindUserByEmail(email)
	if err != nil {
		fmt.Printf("error from model")
		return "", err // return empty token
	}

	// Verify password
	if FoundUser.Password != utils.HashPassword(password, FoundUser.Salt) {
		return "", errors.New("invalid password") // return empty token
	}

	// Generate JWT token
	token, err := generateJWT(FoundUser)
	if err != nil {
		return "", errors.New("failed to generate token")
	}

	return token, nil // Success case: return JWT token
}
