package service

import "github.com/asadbek21coder/fintracker/internal/repository"

type Service struct {
}

func NewService(repos *repository.Repository) *Service {
	return &Service{}
}
