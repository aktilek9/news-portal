package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ErrorResponse(c *gin.Context, err error) {
	if appErr, ok := err.(*AppError); ok {
		c.JSON(appErr.Code, gin.H{
			"message": appErr.Message,
		})
		return
	}
	// TODO: add logs,
	c.JSON(http.StatusInternalServerError, gin.H{
		"message": "unknown",
		"error":   err,
	})
}

func ErrorResponseWithAbort(c *gin.Context, code int, message string) {
	c.AbortWithStatusJSON(code, gin.H{
		"message": message,
	})
}

func SuccessResponse(c *gin.Context, code int, data interface{}) {
	c.JSON(code, data)
}
