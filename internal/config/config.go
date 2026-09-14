package config

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

type Config struct {
	App      AppConfig
	Database DatabaseConfig
}

type AppConfig struct {
	Environment string
	Host        string
	Port        string
	Timezone    string
}

type DatabaseConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
	SSLMode  string
}

func Load() (Config, error) {
	cfg := Config{
		App: AppConfig{
			Environment: getEnv("APP_ENV", "development"),
			Host:        getEnv("APP_HOST", "127.0.0.1"),
			Port:        getEnv("APP_PORT", "8080"),
			Timezone:    getEnv("APP_TIMEZONE", "Asia/Jakarta"),
		},

		Database: DatabaseConfig{
			Host:     getEnv("DB_HOST", "localhost"),
			Port:     getEnv("DB_PORT", "5432"),
			User:     getEnv("DB_USER", "postgres"),
			Password: os.Getenv("DB_PASSWORD"),
			Name:     getEnv("DB_NAME", "rest_api"),
			SSLMode:  getEnv("DB_SSLMODE", "disable"),
		},
	}

	port, err := strconv.Atoi(cfg.App.Port)
	if err != nil || port < 1 || port > 65535 {
		return Config{}, fmt.Errorf("invalid APP_PORT %q", cfg.App.Port)
	}
	if _, err := time.LoadLocation(cfg.App.Timezone); err != nil {
		return Config{}, fmt.Errorf("invalid APP_TIMEZONE %q: %w", cfg.App.Timezone, err)
	}
	switch cfg.Database.SSLMode {
	case "disable", "require", "verify-ca", "verify-full":
	default:
		return Config{}, fmt.Errorf("invalid DB_SSLMODE %q", cfg.Database.SSLMode)
	}

	return cfg, nil
}

func getEnv(key, fallback string) string {
	value := os.Getenv(key)

	if value == "" {
		return fallback
	}

	return value
}
