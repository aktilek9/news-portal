package middleware

import (
	"net/http"
	"news-portal/pkg/response"
	"slices"

	"github.com/gin-gonic/gin"
)

func CheckPermission(allowedRoles []string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		role := ctx.GetHeader(roleCtx)
		if !slices.Contains(allowedRoles, role) {
			response.ErrorResponseWithAbort(ctx, http.StatusForbidden, "You are not allowed")
			return
		}
		ctx.Next()
	}
}
