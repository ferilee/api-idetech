package config

import (
	"log"
	"os"
	"strings"
)

type Config struct {
	AppEnv          string
	Port            string
	BaseURL         string
	AllowedOrigins  []string
	JWTIssuer       string
	JWTAudience     string
	JWTSecret       string
	PostgresHost    string
	PostgresPort    string
	PostgresDB      string
	PostgresUser    string
	PostgresPass    string
	PostgresSSLMode string
}

func MustLoad() Config {
	cfg := Config{
		AppEnv:          getEnv("APP_ENV", "development"),
		Port:            getEnv("APP_PORT", "8080"),
		BaseURL:         getEnv("APP_BASE_URL", "http://localhost:8080"),
		AllowedOrigins:  splitCSV(getEnv("APP_ALLOWED_ORIGINS", "http://localhost:3000")),
		JWTIssuer:       getEnv("JWT_ISSUER", "idetech-api"),
		JWTAudience:     getEnv("JWT_AUDIENCE", "idetech-web"),
		JWTSecret:       getEnv("JWT_SECRET", "change-me"),
		PostgresHost:    getEnv("POSTGRES_HOST", ""),
		PostgresPort:    getEnv("POSTGRES_PORT", "5432"),
		PostgresDB:      getEnv("POSTGRES_DB", "idetech"),
		PostgresUser:    getEnv("POSTGRES_USER", "idetech"),
		PostgresPass:    getEnv("POSTGRES_PASSWORD", "idetech"),
		PostgresSSLMode: getEnv("POSTGRES_SSLMODE", "disable"),
	}

	log.Printf("config loaded env=%s port=%s", cfg.AppEnv, cfg.Port)
	return cfg
}

func (c Config) PostgresDSN() string {
	if c.PostgresHost == "" {
		return ""
	}

	return "postgres://" + c.PostgresUser + ":" + c.PostgresPass + "@" + c.PostgresHost + ":" + c.PostgresPort + "/" + c.PostgresDB + "?sslmode=" + c.PostgresSSLMode
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
