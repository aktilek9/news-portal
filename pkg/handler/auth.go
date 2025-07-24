package handler

import (
	"net/http"
	"news-portal/dto"

	"github.com/gin-gonic/gin"
)

func (h *Handler) Login(c *gin.Context) {
	var body LoginUserRequest
	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})
		return
	}

	token, err := h.service.Login(body.Email, body.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "failed to login",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}

func (h *Handler) Register(c *gin.Context) {
	var body CreateUserRequest
	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})
		return
	}

	userDTO := dto.UserDto{
		Role:     body.Role,
		Email:    body.Email,
		Password: body.Password,
	}

	id, err := h.service.Register(&userDTO)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to register user",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{"id": id})
}

// 1. Get request body,
// 2. convert request body to user dto
// 3. send request to service
