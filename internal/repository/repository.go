package repository

import (
	"github.com/asadbek21coder/fintracker/config"
	"github.com/asadbek21coder/fintracker/internal/logger"
	"github.com/jmoiron/sqlx"
)

type Repository struct {
}

func NewRepository(db *sqlx.DB, log *logger.Logger, cfg *config.Config) *Repository {
	return &Repository{}
}
