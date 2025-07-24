package handler

import (
	"news-portal/pkg/service"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service service.Service
}

func newHandler(service service.Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) TestFunc(c *gin.Context) {
}
