package config

import (
	"context"
	"database/sql"
	"market-suggester/internal/db"
)

func SetupDevelopmentData(cfg *Config,dbConn *sql.DB) {
	if cfg.Environment == "production" {
		return
	}
	ctx := context.Background()
	defer ctx.Done()
	queries := db.New(dbConn)
	queries.CreateUser(ctx, db.CreateUserParams{
		Email: "admin@test.com",
		Name:  "admin",
	})

}