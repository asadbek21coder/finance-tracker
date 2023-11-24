package service

import (
	"github.com/asadbek21coder/fintracker/internal/entities"
	"github.com/asadbek21coder/fintracker/internal/repository"
)

type UserService struct {
	repo repository.User
}

func NewUserService(repo repository.User) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) CreateUser(req entities.CreateUserReq) (entities.User, error) {
	return s.repo.CreateUser(req)
}

func (s *UserService) GetAllUser(limit, offset int) ([]entities.User, error) {
	return s.repo.GetAllUser(limit, offset)
}

func (s *UserService) GetUserByID(id int) (entities.User, error) {
	return s.repo.GetUserByID(id)
}

func (s *UserService) UpdateUser(req entities.UpdateUserReq) (entities.User, error) {
	return s.repo.UpdateUser(req)
}

func (s *UserService) DeleteUser(id int) error {
	return s.repo.DeleteUser(id)
}
