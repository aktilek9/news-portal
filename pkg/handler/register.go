package handler

import (
	"news-portal/pkg/service"

	"github.com/gin-gonic/gin"
)

func RegisterEndpoint(router *gin.Engine, service service.Service) {
	h := newHandler(service)

	authGroup := router.Group("/auth")
	{
		authGroup.POST("/login", h.Login)
		authGroup.POST("/register", h.Register)
	}
}
