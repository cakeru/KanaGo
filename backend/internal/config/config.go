package config

import (
	"os"
	"strconv"
)

type Config struct {
	Server ServerConfig
	DB    DBConfig
	JWT  JWTConfig
	CORS  CORSConfig
}

type ServerConfig struct {
	Port string
	Mode string
}

type DBConfig struct {
	Host	 string
	Port     int
	User     string
	Password string
	DBName   string
}

type JWTConfig struct {
	Secret        string
	ExpiryMinutes int
	RefreshExpiryDays int
}

type CORSConfig struct {
	AllowedOrigins []string
}

func Load() *Config {
	return &Config{
		Server: ServerConfig{
			Port: getEnv("SERVER_PORT", "8080"),
			Mode: getEnv("SERVER_MODE", "development"),
		},
		DB: DBConfig{
			Host:	 getEnv("DB_HOST", "localhost"),
			Port:     getEnvInt("DB_PORT", 5432),
			User:     getEnv("DB_USER", "postgres"),
			Password: getEnv("DB_PASSWORD", "password"),
			DBName:   getEnv("DB_NAME", "kanago_dev"),
		},
		JWT: JWTConfig{
			Secret:        getEnv("JWT_SECRET", "secret"),
			ExpiryMinutes: getEnvInt("JWT_EXPIRY_MINUTES", 15),
			RefreshExpiryDays: getEnvInt("JWT_REFRESH_EXPIRY_DAYS", 7),
		},
		CORS: CORSConfig{
			AllowedOrigins: []string{"http://localhost:5173", "http://localhost:3000"},
		},
	}
}

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}

func getEnvInt(key string, fallback int) int {
	val := getEnv(key, "")
	if val == "" {
		return fallback
	}
	if intVal, err:= strconv.Atoi(val); err == nil {
		return intVal
	}
	return fallback
}