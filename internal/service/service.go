package service

import (
	"github.com/asadbek21coder/fintracker/internal/entities"
	"github.com/asadbek21coder/fintracker/internal/repository"
)

type Service struct {
	User
}

type User interface {
	CreateUser(entities.CreateUserReq) (entities.User, error)
	GetAllUser(limit, offset int) ([]entities.User, error)
	GetUserByID(id int) (entities.User, error)
	UpdateUser(entities.UpdateUserReq) (entities.User, error)
	DeleteUser(id int) error
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		User: NewUserService(repos.User),
	}
}
