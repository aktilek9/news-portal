package handler

import (
	"news-portal/pkg/service"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	Service service.Service
}

func NewHandler(service service.Service) *Handler {
	return &Handler{Service: service}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	return router
}
