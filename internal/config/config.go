package config

import (
	"github.com/kelseyhightower/envconfig"
	"go.uber.org/zap"
)

type Config struct {
	Port              string          `envconfig:"PORT" default:":8080"`
	DatabaseName      string          `envconfig:"DATABASE_NAME" default:"postgres"`
	DatabaseUser      string          `envconfig:"DATABASE_USER" default:"postgres"`
	DatabasePassword  string          `envconfig:"DATABASE_PASSWORD" default:"postgres"`
	LogLevel          zap.AtomicLevel `envconfig:"LOG_LEVEL" default:"debug"`
	Environment       string          `envconfig:"ENV" default:"development"`
}

func loadConfig() (*Config, error) {
	var cfg Config
	err := envconfig.Process("", &cfg)
	if err != nil {
		return nil, err
	}
	return &cfg, nil
}

func MustLoadConfig() *Config {
	cfg, err := loadConfig()
	if err != nil {
		panic(err)
	}
	return cfg
}
