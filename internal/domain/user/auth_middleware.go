package user

import (
	"context"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/msoft-g1/todo-list-backend/internal/errs"
)

var userCtxKey = &contextKey{"user"}

type contextKey struct {
	name string
}

// AuthMiddleware defines Auth Middleware
func AuthMiddleware(userService *Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			// continue as non-authenticated
			return
		}
		authHeaderParts := strings.Fields(authHeader)
		if len(authHeaderParts) != 2 || strings.ToLower(authHeaderParts[0]) != "bearer" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, errs.New(errs.CodeInvalidTokenFormat, "Formato de token inválido"))
			return
		}
		data, err := userService.ValidateAccessToken(authHeaderParts[1])
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, errs.New(errs.CodeInvalidToken, "Token inválido"))
			return
		}
		// --- Pass the token data object through the request context
		ctx := context.WithValue(c.Request.Context(), userCtxKey, data)
		c.Request = c.Request.WithContext(ctx)
	}
}

// AuthForContext returns the auth data for the specified context.
func AuthForContext(ctx context.Context) *AccessTokenData {
	d, ok := ctx.Value(userCtxKey).(*AccessTokenData)
	if !ok {
		return nil
	}
	return d
}
