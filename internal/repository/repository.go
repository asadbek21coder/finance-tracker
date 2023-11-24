package repository

import (
	"github.com/asadbek21coder/fintracker/config"
	"github.com/asadbek21coder/fintracker/internal/entities"
	"github.com/asadbek21coder/fintracker/internal/logger"
	"github.com/jmoiron/sqlx"
)

type Repository struct {
	User
}

type User interface {
	CreateUser(entities.CreateUserReq) (entities.User, error)
	GetAllUser(limit, offset int) ([]entities.User, error)
	GetUserByID(id int) (entities.User, error)
	UpdateUser(entities.UpdateUserReq) (entities.User, error)
	DeleteUser(id int) error
}

func NewRepository(db *sqlx.DB, log *logger.Logger, cfg *config.Config) *Repository {
	return &Repository{
		User: NewUserPostgres(db, log, cfg),
	}
}
