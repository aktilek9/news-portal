package handler

import (
	"net/http"
	"news-portal/dto"
	"news-portal/middleware"
	"news-portal/pkg/response"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) CreateComment(c *gin.Context) {
	var body CreateCommentRequest
	if err := c.BindJSON(&body); err != nil {
		response.ErrorResponse(c, response.NewAppError(http.StatusBadRequest, "Failed to read body", err))
		return
	}

	authorID, err := middleware.GetUserID(c)
	if err != nil {
		response.ErrorResponse(c, response.NewAppError(http.StatusForbidden, "Cannot find user id", err))
		return
	}

	commentDTO := dto.CommentDTO{
		Content:  body.Content,
		AuthorID: authorID,
		NewsID:   body.NewsID,
	}

	if err := h.service.CreateComment(commentDTO); err != nil {
		response.ErrorResponse(c, response.NewAppError(http.StatusInternalServerError, "Failed to create comment", err))
		return
	}
	response.SuccessResponse(c, http.StatusOK, gin.H{
		"message": "comment created",
	})
}

func (h *Handler) GetCommentsByNewsID(c *gin.Context) {
	idStr := c.Param("id")
	newsID, err := strconv.Atoi(idStr)
	if err != nil {
		response.ErrorResponse(c, response.NewAppError(http.StatusBadRequest, "Invalid news id", err))
		return
	}

	comments, err := h.service.GetCommentsByNewsID(newsID)
	if err != nil {
		response.ErrorResponse(c, response.NewAppError(http.StatusInternalServerError, "Failed to get comments", err))
		return
	}
	response.SuccessResponse(c, http.StatusOK, comments)
}

func (h *Handler) DeleteComment(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		response.ErrorResponse(c, response.NewAppError(http.StatusBadRequest, "Invalid ID", err))
		return
	}

	userID, err := middleware.GetUserID(c)
	if err != nil {
		response.ErrorResponse(c, response.NewAppError(http.StatusForbidden, "Cannot find user id", err))
		return
	}

	userRole := middleware.GetUserRole(c)

	// Author verification
	comment, err := h.service.GetCommentByID(id)
	if err != nil {
		response.ErrorResponse(c, response.NewAppError(http.StatusInternalServerError, "Comment not found", err))
		return
	}

	if userRole != "admin" {
		if comment.AuthorID != uint(userID) {
			response.ErrorResponse(c, response.NewAppError(http.StatusForbidden, "You are not the author of this comment", err))
			return
		}
	}

	// Deleting a comment
	if err := h.service.DeleteComment(id); err != nil {
		response.ErrorResponse(c, response.NewAppError(http.StatusInternalServerError, "Failed to delete the comment", err))
		return
	}

	response.SuccessResponse(c, http.StatusOK, gin.H{
		"message": "Comment deleted",
	})
}
