package service

import (
	"github.com/c4miloarriagada/keys-be/internal/core/entity"
)

type UserRepository interface {
	GetAll() ([]entity.User, error)
}

type UserService struct {
	repo UserRepository
}

func NewUserService(repo UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) GetUsers() ([]entity.User, error) {
	return s.repo.GetAll()
}
