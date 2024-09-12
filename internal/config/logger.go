package config

import (
	"fmt"

	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func SetupLogger(cfg *Config) (*zap.Logger) {
	var logger *zap.Logger
	var err error
	if cfg.Environment == "production" {
		logger, err = zap.NewProduction()
	} else {
		logger, err = zap.NewDevelopment()
	}
	if err != nil {
		panic(fmt.Sprintf("Failed to initialize logger: %v", err))
	}
	defer logger.Sync()
	return logger
}

func ZapMiddleware(logger *zap.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		query := c.Request.URL.RawQuery

		c.Next()

		end := time.Now()
		latency := end.Sub(start)
		status := c.Writer.Status()

		logFields := []zap.Field{
			zap.Int("status", status),
			zap.String("method", c.Request.Method),
			zap.String("path", path),
			zap.String("query", query),
			zap.Duration("latency", latency),
			zap.String("ip", c.ClientIP()),
			zap.String("user-agent", c.Request.UserAgent()),
		}

		switch {
		case status >= 500:
			logger.Error("Server Error", logFields...)
		case status >= 400:
			logger.Warn("Client Error", logFields...)
		case status >= 300:
			logger.Info("Redirection", logFields...)
		default:
			logger.Info("Success", logFields...)
		}
	}
}