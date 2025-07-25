package handler

import (
	"news-portal/pkg/jwt"
	"news-portal/pkg/service"

	"github.com/gin-gonic/gin"
)

func RegisterEndpoint(router *gin.Engine, service service.Service, jwt jwt.JWTService) {
	h := newHandler(service)

	authGroup := router.Group("/auth")
	{
		// authGroup.Use(middleware.UserIdentify(jwt))
		authGroup.POST("/login", h.Login)
		authGroup.POST("/register", h.Register)
	}
}
