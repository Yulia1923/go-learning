package service

import (
	"errors"
	"go-learning/day26/model"
	"go-learning/day26/repository"
)

type UserService struct {
	userRepo *repository.UserRepository
}

func NewUserService(userRepo *repository.UserRepository) *UserService {
	return &UserService{
		userRepo: userRepo,
	}
}

func (s *UserService) GetUserByID(id uint) (*model.User, error) {
	if id == 0 {
		return nil, errors.New("用户ID必须大于0")
	}

	return s.userRepo.GetByID(id)
}
