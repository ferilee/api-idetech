package config

import (
	"log"
	"os"
	"strings"
)

type Config struct {
	AppEnv         string
	Port           string
	BaseURL        string
	AllowedOrigins []string
}

func MustLoad() Config {
	cfg := Config{
		AppEnv:         getEnv("APP_ENV", "development"),
		Port:           getEnv("APP_PORT", "8080"),
		BaseURL:        getEnv("APP_BASE_URL", "http://localhost:8080"),
		AllowedOrigins: splitCSV(getEnv("APP_ALLOWED_ORIGINS", "http://localhost:3000")),
	}

	log.Printf("config loaded env=%s port=%s", cfg.AppEnv, cfg.Port)
	return cfg
}

func getEnv(key, fallback string) string {
	value := strings.TrimSpace(os.Getenv(key))
	if value == "" {
		return fallback
	}
	return value
}

func splitCSV(value string) []string {
	raw := strings.Split(value, ",")
	items := make([]string, 0, len(raw))
	for _, item := range raw {
		item = strings.TrimSpace(item)
		if item != "" {
			items = append(items, item)
		}
	}
	return items
}
