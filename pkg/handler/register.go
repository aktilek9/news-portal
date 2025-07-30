package handler

import (
	"news-portal/middleware"
	"news-portal/pkg/jwt"
	"news-portal/pkg/service"

	"github.com/gin-gonic/gin"
)

func RegisterEndpoint(router *gin.Engine, service service.Service, jwt jwt.JWTService) {
	h := newHandler(service)

	authGroup := router.Group("/auth")
	{
		authGroup.POST("/login", h.Login)
		authGroup.POST("/register", h.Register)
	}

	newsGroup := router.Group("/news")
	newsGroup.Use(middleware.UserIdentify(jwt))
	{
		newsGroup.GET("/", h.GetAllNews)
		newsGroup.GET("/:id", h.GetNewsByID)  
		newsGroup.POST("/",middleware.CheckPermission([]string{"admin", "manager"}), h.CreateNews)      // Only manager 
		newsGroup.PUT("/:id", h.UpdateNews)    // Only the author of the news
		newsGroup.DELETE("/:id", h.DeleteNews) // Only the author of the news
	}
}