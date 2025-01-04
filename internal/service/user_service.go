package service

import "github.com/c4miloarriagada/keys-be/internal/domain"

type UserService struct {
	UserRepo domain.UserRepository
}

func NewUserService(repo domain.UserRepository) *UserService {
	return &UserService{UserRepo: repo}
}

func (s *UserService) GetUserByID(id int) (*domain.User, error) {
	return s.UserRepo.GetByID(id)
}

func (s *UserService) CreateUser(user *domain.User) error {
	return s.UserRepo.Save(user)
}

func (s *UserService) GetAllUsers() ([]domain.User, error) {
	return s.UserRepo.GetAll()
}
