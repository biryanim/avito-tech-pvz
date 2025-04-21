package middleware

import (
	"github.com/biryanim/avito-tech-pvz/pkg/logger"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"time"
)

func LogMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		query := c.Request.URL.RawQuery

		c.Next()

		end := time.Now()
		latency := end.Sub(start)
		status := c.Writer.Status()

		fields := []zap.Field{
			zap.Int("status", status),
			zap.String("method", c.Request.Method),
			zap.String("query", query),
			zap.String("ip", c.ClientIP()),
			zap.String("user-agent", c.Request.UserAgent()),
			zap.Duration("latency", latency),
		}

		switch {
		case status >= 500:
			logger.Error(path, fields...)
		case status >= 400:
			logger.Warn(path, fields...)
		case len(c.Errors) > 0:
			for _, e := range c.Errors {
				logger.Error(path,
					append(fields, zap.String("error", e.Error()))...)
			}
		default:
			logger.Info(path, fields...)
		}
	}
}
