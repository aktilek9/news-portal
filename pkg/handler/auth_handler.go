package handler

import (
	"net/http"
	"news-portal/dto"
	"news-portal/pkg/response"

	"github.com/gin-gonic/gin"
)

func (h *Handler) Login(c *gin.Context) {
	var body LoginUserRequest
	if err := c.BindJSON(&body); err != nil {
		response.ErrorResponse(c, response.NewAppError(http.StatusBadRequest, "Failed to read body", err))
		return
	}

	token, err := h.service.Login(body.Email, body.Password)
	if err != nil {
		response.ErrorResponse(c, response.NewAppError(http.StatusBadRequest, err.Error(), err))
		return
	}

	response.SuccessResponse(c, http.StatusOK, gin.H{
		"token": token,
	})
}

func (h *Handler) Register(c *gin.Context) {
	var body CreateUserRequest
	if err := c.BindJSON(&body); err != nil {
		response.ErrorResponse(c, response.NewAppError(http.StatusBadRequest, "Failed to read body", err))
		return
	}

	userDTO := dto.UserDto{
		Email:    body.Email,
		Password: body.Password,
	}

	id, err := h.service.Register(&userDTO)
	if err != nil {
		response.ErrorResponse(c, response.NewAppError(http.StatusBadRequest, "Failed to register user", err))
		return
	}

	response.SuccessResponse(c, http.StatusOK, gin.H{
		"id": id,
	})
}
