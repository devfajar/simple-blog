package service

import (
	"errors"

	"github.com/devfajar/blog-app/internal/config"
	"github.com/devfajar/blog-app/internal/repository"
	"github.com/devfajar/blog-app/internal/util"
)

type AuthService interface {
	Login(email, password string) (string, error)
}

type authService struct {
	users repository.UserRepository
	cfg   config.Config
}

func NewAuthService(users repository.UserRepository, cfg config.Config) AuthService {
	return &authService{users: users, cfg: cfg}
}

func (s *authService) Login(email, password string) (string, error) {
	u, err := s.users.FindByEmail(email)
	if err != nil {
		return "", errors.New("invalid credentials")
	}
	if !util.CheckPasswordHash(password, u.PasswordHash) {
		return "", errors.New("invalid credentials")
	}
	return util.GenerateJWT(s.cfg.JWTSecret, u.ID, u.Role, u.Email, u.Name)
}
