package rest

import (
	"github.com/asadbek21coder/fintracker/internal/logger"
	"github.com/asadbek21coder/fintracker/internal/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
	log      *logger.Logger
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	router.GET("/ping", h.pong)

	return router
}
