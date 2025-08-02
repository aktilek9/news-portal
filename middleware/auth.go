package middleware

import (
	"net/http"
	"news-portal/pkg/jwt"
	"news-portal/pkg/response"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

const (
	authorizationHeader string = "Authorization"
	userCtx             string = "userId"
	roleCtx             string = "role"
)

func UserIdentify(jwt jwt.JWTService) gin.HandlerFunc {
	return func(c *gin.Context) {
		header := c.GetHeader(authorizationHeader)
		if header == "" {
			response.ErrorResponseWithAbort(c, http.StatusUnauthorized, "empty auth header")
			return
		}

		parts := strings.Split(header, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			response.ErrorResponseWithAbort(c, http.StatusUnauthorized, "invalid auth header")
			return
		}

		tokenStr := parts[1]
		if tokenStr == "" {
			response.ErrorResponseWithAbort(c, http.StatusUnauthorized, "token is empty")
			return
		}

		userID, role, err := jwt.ParseToken(tokenStr)
		if err != nil {
			response.ErrorResponse(c, err)
			c.Abort()
			return
		}
		c.Request.Header.Add(userCtx, strconv.Itoa(userID))
		c.Request.Header.Add(roleCtx, role)

		c.Next()
	}
}

func GetUserID(c *gin.Context) (int, error) {
	userID, err := strconv.Atoi(c.GetHeader(userCtx))
	if err != nil {
		return -1, err
	}
	return userID, nil
}

func GetUserRole(c *gin.Context) string {
	userRole := c.GetHeader(roleCtx)
	return userRole
}
