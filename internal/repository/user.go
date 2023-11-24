package repository

import (
	"github.com/asadbek21coder/fintracker/config"
	"github.com/asadbek21coder/fintracker/internal/entities"
	"github.com/asadbek21coder/fintracker/internal/logger"
	"github.com/jmoiron/sqlx"
)

type UserPostgres struct {
	Log *logger.Logger
	Cfg *config.Config
	db  *sqlx.DB
}

func NewUserPostgres(db *sqlx.DB, log *logger.Logger, cfg *config.Config) *UserPostgres {

	return &UserPostgres{}
}

func (u *UserPostgres) CreateUser(entities.CreateUserReq) (entities.User, error) {
	userRes := entities.User{}

	return userRes, nil
}

func (u *UserPostgres) GetAllUser(limit, offset int) ([]entities.User, error) {
	userRes := []entities.User{}

	return userRes, nil
}

func (u *UserPostgres) GetUserByID(id int) (entities.User, error) {
	userRes := entities.User{}

	return userRes, nil
}

func (u *UserPostgres) UpdateUser(req entities.UpdateUserReq) (entities.User, error) {
	userRes := entities.User{}

	return userRes, nil
}

func (u *UserPostgres) DeleteUser(id int) error {

	return nil
}
