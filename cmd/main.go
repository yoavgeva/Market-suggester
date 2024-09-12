package main

import (

	"fmt"


	"market-suggester/internal/config"
	"market-suggester/internal/handlers"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

)
var logger *zap.Logger

func main() {
	cfg := config.MustLoadConfig()
	ginServer := setupRouter(cfg)
	if err := config.RunMigrations(cfg); err != nil {
		panic(err)
	}
	setupGinMiddleware(ginServer,cfg)
	ginServer.GET("/health", handlers.HealthHandler)

	config.SetupDevelopmentData(cfg)

	err := ginServer.Run(":8080")
	if err != nil {
		logger.Error("Failed to start server",
			zap.Error(err),
			zap.String("port", cfg.Port),
		)
		panic(err)
	}
}

func setupRouter(cfg *config.Config) *gin.Engine {
	var r *gin.Engine
	
	if cfg.Environment == "production" {
		gin.SetMode(gin.ReleaseMode)
		r = gin.New()
	} else {
		r = gin.Default()
	}
	return r
}

func setupGinMiddleware(r *gin.Engine,cfg *config.Config) {
	r.Use(config.ZapMiddleware(config.SetupLogger(cfg)))
	r.Use(gin.Recovery())
}


