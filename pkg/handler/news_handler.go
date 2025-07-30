package handler

import (
	"errors"
	"net/http"
	"news-portal/dto"
	"news-portal/middleware"
	"news-portal/pkg/response"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func (h *Handler) GetAllNews(c *gin.Context) {
	news, err := h.service.GetAllNews()
	if err != nil {
		response.ErrorResponse(c, response.NewAppError(http.StatusInternalServerError, "Failed to get news", err))
		return
	}
	response.SuccessResponse(c, http.StatusOK, news)
}

func (h *Handler) GetNewsByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		response.ErrorResponse(c, response.NewAppError(http.StatusBadRequest, "Invalid ID", err))
		return
	}

	news, err := h.service.GetNewsByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			response.ErrorResponse(c, response.NewAppError(http.StatusNotFound, "News not found", err))
		} else {
			response.ErrorResponse(c, response.NewAppError(http.StatusInternalServerError, "Failed to get news", err))
		}
		return
	}
	response.SuccessResponse(c, http.StatusOK, news)
}

func (h *Handler) CreateNews(c *gin.Context) {
	var body CreateNewsRequest
	if err := c.BindJSON(&body); err != nil {
		response.ErrorResponse(c, response.NewAppError(http.StatusBadRequest, "Failed to read body", err))
		return
	}

	authorID, err := middleware.GetUserID(c)
	if err != nil {
		response.ErrorResponse(c, response.NewAppError(http.StatusForbidden, "Cannot find user id", err))
		return
	}

	news := dto.News{
		Title:   body.Title,
		Content: body.Content,
		AuthorID: authorID,
	}

	if err := h.service.CreateNews(&news); err != nil {
		response.ErrorResponse(c, response.NewAppError(http.StatusInternalServerError, "Failed to create news", err))
		return
	}
	response.SuccessResponse(c, http.StatusCreated, gin.H{
		"message": "News created succesfully",
	})
}

func (h *Handler) UpdateNews(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		response.ErrorResponse(c, response.NewAppError(http.StatusBadRequest, "Invalid ID", err))
		return
	}

	var news dto.News
	if err := c.BindJSON(&news); err != nil {
		response.ErrorResponse(c, response.NewAppError(http.StatusBadRequest, "Failed to read body", err))
		return
	}

	if err := h.service.UpdateNews(id, news); err != nil {
		response.ErrorResponse(c, response.NewAppError(http.StatusInternalServerError, "Failed to update news", err))
		return
	}
	response.SuccessResponse(c, http.StatusOK, gin.H{
		"message": "News updated",
	})
}

func (h *Handler) DeleteNews(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		response.ErrorResponse(c, response.NewAppError(http.StatusBadRequest, "Invalid ID", err))
		return
	}

	if err := h.service.DeleteNews(id); err != nil {
		response.ErrorResponse(c, response.NewAppError(http.StatusNotFound, "News not found", err))
		return
	}
	response.SuccessResponse(c, http.StatusOK, gin.H{
		"message": "News deleted",
	})
}
