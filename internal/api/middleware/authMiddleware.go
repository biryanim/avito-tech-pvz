package middleware

import (
	"github.com/biryanim/avito-tech-pvz/internal/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

const (
	auth = "Bearer "
)

func AuthMiddleware(authService service.AuthService) gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code":    http.StatusUnauthorized,
				"message": "token required",
			})
			c.Abort()
			return
		}

		token = strings.TrimPrefix(token, auth)
		access, err := authService.Check(c.Request.Context(), token, c.Request.Method, c.FullPath())
		if err != nil || !access {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code":    http.StatusUnauthorized,
				"message": "token invalid",
			})
			c.Abort()
			return
		}

		c.Next()
	}
}
