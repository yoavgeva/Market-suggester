package main

import (
	"database/sql"
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
		logger.Fatal("Failed to do migrations:", zap.Error(err))
		panic(err)
	}
	dbConn, err := sql.Open("postgres", fmt.Sprintf("postgresql://%s:%s@localhost:5432/%s?sslmode=disable", cfg.DatabaseUser, cfg.DatabasePassword, cfg.DatabaseName))
	if err != nil {
		logger.Fatal("Failed to connect to database:", zap.Error(err))
		panic(err)
	}
	defer dbConn.Close()
	
	setupGinMiddleware(ginServer,cfg)
	ginServer.GET("/health", handlers.HealthHandler)

	config.SetupDevelopmentData(cfg,dbConn)

	err = ginServer.Run(":8080")
	if err != nil {
		logger.Fatal("Failed to start server",
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


