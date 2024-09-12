package config

import "market-suggester/internal/db"

func SetupDevelopmentData(cfg *Config) {
	if cfg.Environment == "production" {
		return
	}
	



}