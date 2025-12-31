# Backend Setup Guide: Building with Go

This guide walks you through setting up a Go backend for your Japanese learning platform.

---

## Prerequisites

- **Go 1.21+** - Download from https://golang.org/
- **PostgreSQL 14+** - Database server
- **Git** - Version control
- **Postman or Thunder Client** - API testing (optional but recommended)

---

## Project Initialization

### Step 1: Create Project Directory

```bash
mkdir go-kanadojo-backend
cd go-kanadojo-backend
```

### Step 2: Initialize Go Module

```bash
go mod init github.com/yourusername/go-kanadojo-backend
```

### Step 3: Create Directory Structure

```bash
# Create main directories
mkdir -p cmd/server
mkdir -p internal/{api,models,services,repository,db,config,logger,auth,middleware}
mkdir -p pkg/{utils,constants}
mkdir -p migrations
mkdir -p data
mkdir -p tests

# Create initial files
touch cmd/server/main.go
touch internal/api/routes.go
touch internal/api/response.go
touch internal/config/config.go
touch internal/db/db.go
touch .env.example
touch docker-compose.yml
touch Dockerfile
touch Makefile
```

---

## Core Dependencies

### Install Required Packages

```bash
# Web framework
go get -u github.com/gin-gonic/gin

# Database
go get -u github.com/lib/pq
go get -u github.com/jmoiron/sqlx

# Migrations
go get -u github.com/golang-migrate/migrate/v4
go get -u github.com/golang-migrate/migrate/v4/database/postgres

# Validation
go get -u github.com/go-playground/validator/v10

# JWT
go get -u github.com/golang-jwt/jwt/v5

# Password hashing
go get -u golang.org/x/crypto

# CORS
go get -u github.com/gin-contrib/cors

# Environment variables
go get -u github.com/joho/godotenv

# Logging
go get -u go.uber.org/zap

# UUID
go get -u github.com/google/uuid

# Timestamps
go get -u github.com/lib/pq

# Optional: Hot reload for development
go get -u github.com/cosmtrek/air
```

---

## Configuration Setup

### `internal/config/config.go`

```go
package config

import (
	"fmt"
	"os"
	"strconv"
	"sync"
)

type Config struct {
	Server   ServerConfig
	Database DatabaseConfig
	JWT      JWTConfig
	CORS     CORSConfig
}

type ServerConfig struct {
	Port string
	Env  string // "development" or "production"
}

type DatabaseConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
	SSLMode  string
}

type JWTConfig struct {
	SecretKey            string
	AccessTokenExpiry    int // in minutes
	RefreshTokenExpiry   int // in days
}

type CORSConfig struct {
	AllowedOrigins []string
	AllowedMethods []string
	AllowedHeaders []string
}

var (
	config *Config
	once   sync.Once
)

// Load configuration from environment variables
func Load() *Config {
	once.Do(func() {
		config = &Config{
			Server: ServerConfig{
				Port: getEnv("SERVER_PORT", "8080"),
				Env:  getEnv("ENV", "development"),
			},
			Database: DatabaseConfig{
				Host:     getEnv("DB_HOST", "localhost"),
				Port:     getEnv("DB_PORT", "5432"),
				User:     getEnv("DB_USER", "postgres"),
				Password: getEnv("DB_PASSWORD", ""),
				DBName:   getEnv("DB_NAME", "kanadojo"),
				SSLMode:  getEnv("DB_SSLMODE", "disable"),
			},
			JWT: JWTConfig{
				SecretKey:          getEnv("JWT_SECRET", "your-secret-key"),
				AccessTokenExpiry:  getEnvInt("JWT_ACCESS_EXPIRY", 15), // 15 minutes
				RefreshTokenExpiry: getEnvInt("JWT_REFRESH_EXPIRY", 7), // 7 days
			},
			CORS: CORSConfig{
				AllowedOrigins: []string{"http://localhost:3000", "http://localhost:5173"},
				AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
				AllowedHeaders: []string{"Content-Type", "Authorization"},
			},
		}
	})
	return config
}

// Get configuration (must call Load() first)
func Get() *Config {
	if config == nil {
		return Load()
	}
	return config
}

func (c *DatabaseConfig) DSN() string {
	return fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		c.Host, c.Port, c.User, c.Password, c.DBName, c.SSLMode,
	)
}

func getEnv(key, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultVal
}

func getEnvInt(key string, defaultVal int) int {
	if value, exists := os.LookupEnv(key); exists {
		if intVal, err := strconv.Atoi(value); err == nil {
			return intVal
		}
	}
	return defaultVal
}
```

### `.env.example`

```env
# Server
SERVER_PORT=8080
ENV=development

# Database
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=your_password
DB_NAME=kanadojo
DB_SSLMODE=disable

# JWT
JWT_SECRET=your-super-secret-jwt-key-change-this
JWT_ACCESS_EXPIRY=15
JWT_REFRESH_EXPIRY=7

# CORS Origins
CORS_ALLOWED_ORIGINS=http://localhost:3000,http://localhost:5173
```

---

## Database Setup

### `internal/db/db.go`

```go
package db

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var DB *sqlx.DB

func Init(dsn string) error {
	var err error
	DB, err = sqlx.Open("postgres", dsn)
	if err != nil {
		return fmt.Errorf("failed to open database: %w", err)
	}

	// Test connection
	if err = DB.Ping(); err != nil {
		return fmt.Errorf("failed to ping database: %w", err)
	}

	// Set connection pool settings
	DB.SetMaxOpenConns(25)
	DB.SetMaxIdleConns(5)

	log.Println("✓ Database connected successfully")
	return nil
}

func Close() error {
	if DB != nil {
		return DB.Close()
	}
	return nil
}
```

### `migrations/001_init_schema.up.sql`

```sql
-- Users table
CREATE TABLE IF NOT EXISTS users (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    email VARCHAR(255) UNIQUE NOT NULL,
    password_hash VARCHAR(255) NOT NULL,
    username VARCHAR(100) UNIQUE NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    email_verified BOOLEAN DEFAULT FALSE,
    last_login TIMESTAMP
);

CREATE INDEX idx_users_email ON users(email);
CREATE INDEX idx_users_username ON users(username);

-- Statistics table
CREATE TABLE IF NOT EXISTS statistics (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID NOT NULL UNIQUE REFERENCES users(id) ON DELETE CASCADE,
    total_answers INT DEFAULT 0,
    correct_answers INT DEFAULT 0,
    accuracy FLOAT DEFAULT 0,
    current_streak INT DEFAULT 0,
    longest_streak INT DEFAULT 0,
    last_practice_date DATE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_statistics_user_id ON statistics(user_id);

-- Preferences table
CREATE TABLE IF NOT EXISTS preferences (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID NOT NULL UNIQUE REFERENCES users(id) ON DELETE CASCADE,
    theme_name VARCHAR(100) DEFAULT 'default',
    font_name VARCHAR(100) DEFAULT 'default',
    sound_enabled BOOLEAN DEFAULT TRUE,
    hotkeys_enabled BOOLEAN DEFAULT TRUE,
    dark_mode BOOLEAN DEFAULT FALSE,
    language VARCHAR(10) DEFAULT 'en',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_preferences_user_id ON preferences(user_id);

-- Progress table (per-content tracking)
CREATE TABLE IF NOT EXISTS progress (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    content_type VARCHAR(50) NOT NULL,
    content_id VARCHAR(100) NOT NULL,
    correct_count INT DEFAULT 0,
    wrong_count INT DEFAULT 0,
    last_practiced TIMESTAMP,
    mastery_level INT DEFAULT 0,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    UNIQUE(user_id, content_type, content_id)
);

CREATE INDEX idx_progress_user_id ON progress(user_id);
CREATE INDEX idx_progress_content ON progress(content_type, content_id);

-- Achievements table
CREATE TABLE IF NOT EXISTS achievements (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(255) UNIQUE NOT NULL,
    description TEXT,
    icon_url VARCHAR(255),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- User achievements (unlocked)
CREATE TABLE IF NOT EXISTS user_achievements (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    achievement_id UUID NOT NULL REFERENCES achievements(id) ON DELETE CASCADE,
    unlocked_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    UNIQUE(user_id, achievement_id)
);

CREATE INDEX idx_user_achievements_user_id ON user_achievements(user_id);
CREATE INDEX idx_user_achievements_achievement_id ON user_achievements(achievement_id);

-- Practice sessions table
CREATE TABLE IF NOT EXISTS practice_sessions (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    content_type VARCHAR(50) NOT NULL,
    game_mode VARCHAR(50) NOT NULL,
    started_at TIMESTAMP NOT NULL,
    ended_at TIMESTAMP,
    correct_answers INT DEFAULT 0,
    total_answers INT DEFAULT 0,
    duration_seconds INT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_practice_sessions_user_id ON practice_sessions(user_id);
CREATE INDEX idx_practice_sessions_started_at ON practice_sessions(started_at);

-- Answer history
CREATE TABLE IF NOT EXISTS answers (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    session_id UUID NOT NULL REFERENCES practice_sessions(id) ON DELETE CASCADE,
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    content_id VARCHAR(100) NOT NULL,
    user_answer VARCHAR(255),
    correct_answer VARCHAR(255) NOT NULL,
    is_correct BOOLEAN NOT NULL,
    time_spent_ms INT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_answers_session_id ON answers(session_id);
CREATE INDEX idx_answers_user_id ON answers(user_id);
```

### `migrations/001_init_schema.down.sql`

```sql
DROP TABLE IF EXISTS answers;
DROP TABLE IF EXISTS practice_sessions;
DROP TABLE IF EXISTS user_achievements;
DROP TABLE IF EXISTS achievements;
DROP TABLE IF EXISTS progress;
DROP TABLE IF EXISTS preferences;
DROP TABLE IF EXISTS statistics;
DROP TABLE IF EXISTS users;
```

---

## Data Models

### `internal/models/user.go`

```go
package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID            uuid.UUID  `db:"id" json:"id"`
	Email         string     `db:"email" json:"email"`
	PasswordHash  string     `db:"password_hash" json:"-"`
	Username      string     `db:"username" json:"username"`
	EmailVerified bool       `db:"email_verified" json:"email_verified"`
	LastLogin     *time.Time `db:"last_login" json:"last_login"`
	CreatedAt     time.Time  `db:"created_at" json:"created_at"`
	UpdatedAt     time.Time  `db:"updated_at" json:"updated_at"`
}

type CreateUserRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Username string `json:"username" binding:"required,min=3,max=100"`
	Password string `json:"password" binding:"required,min=8"`
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type AuthResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	User         User   `json:"user"`
}
```

### `internal/models/progress.go`

```go
package models

import (
	"time"

	"github.com/google/uuid"
)

type Progress struct {
	ID          uuid.UUID `db:"id" json:"id"`
	UserID      uuid.UUID `db:"user_id" json:"user_id"`
	ContentType string    `db:"content_type" json:"content_type"` // "kana", "kanji", "vocabulary"
	ContentID   string    `db:"content_id" json:"content_id"`
	CorrectCount int      `db:"correct_count" json:"correct_count"`
	WrongCount  int       `db:"wrong_count" json:"wrong_count"`
	LastPracticed *time.Time `db:"last_practiced" json:"last_practiced"`
	MasteryLevel int      `db:"mastery_level" json:"mastery_level"`
	CreatedAt   time.Time `db:"created_at" json:"created_at"`
	UpdatedAt   time.Time `db:"updated_at" json:"updated_at"`
}

type UpdateProgressRequest struct {
	ContentType string `json:"content_type" binding:"required"`
	ContentID   string `json:"content_id" binding:"required"`
	IsCorrect   bool   `json:"is_correct"`
}

type ProgressResponse struct {
	Progress      Progress `json:"progress"`
	Accuracy      float64  `json:"accuracy"`
	MasteryLevel  int      `json:"mastery_level"`
}
```

---

## Authentication

### `internal/auth/jwt.go`

```go
package auth

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"yourmodule/internal/config"
)

type Claims struct {
	UserID   uuid.UUID `json:"user_id"`
	Email    string    `json:"email"`
	Username string    `json:"username"`
	jwt.RegisteredClaims
}

func GenerateAccessToken(userID uuid.UUID, email, username string) (string, error) {
	cfg := config.Get()
	expiryTime := time.Now().Add(time.Duration(cfg.JWT.AccessTokenExpiry) * time.Minute)

	claims := &Claims{
		UserID:   userID,
		Email:    email,
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expiryTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(cfg.JWT.SecretKey))
}

func GenerateRefreshToken(userID uuid.UUID) (string, error) {
	cfg := config.Get()
	expiryTime := time.Now().AddDate(0, 0, cfg.JWT.RefreshTokenExpiry)

	claims := jwt.RegisteredClaims{
		Subject:   userID.String(),
		ExpiresAt: jwt.NewNumericDate(expiryTime),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(cfg.JWT.SecretKey))
}

func VerifyToken(tokenString string) (*Claims, error) {
	cfg := config.Get()
	claims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(cfg.JWT.SecretKey), nil
	})

	if err != nil || !token.Valid {
		return nil, fmt.Errorf("invalid token: %w", err)
	}

	return claims, nil
}
```

### `internal/middleware/auth.go`

```go
package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"yourmodule/internal/auth"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "missing authorization header"})
			c.Abort()
			return
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid authorization header"})
			c.Abort()
			return
		}

		claims, err := auth.VerifyToken(parts[1])
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
			c.Abort()
			return
		}

		c.Set("user_id", claims.UserID.String())
		c.Set("email", claims.Email)
		c.Set("username", claims.Username)
		c.Next()
	}
}
```

---

## API Response Format

### `internal/api/response.go`

```go
package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
	Message string      `json:"message,omitempty"`
}

func SuccessResponse(c *gin.Context, statusCode int, data interface{}, message string) {
	c.JSON(statusCode, Response{
		Success: true,
		Data:    data,
		Message: message,
	})
}

func ErrorResponse(c *gin.Context, statusCode int, error string) {
	c.JSON(statusCode, Response{
		Success: false,
		Error:   error,
	})
}

func CreatedResponse(c *gin.Context, data interface{}, message string) {
	SuccessResponse(c, http.StatusCreated, data, message)
}

func OKResponse(c *gin.Context, data interface{}, message string) {
	SuccessResponse(c, http.StatusOK, data, message)
}

func BadRequestResponse(c *gin.Context, error string) {
	ErrorResponse(c, http.StatusBadRequest, error)
}

func UnauthorizedResponse(c *gin.Context) {
	ErrorResponse(c, http.StatusUnauthorized, "unauthorized")
}

func ForbiddenResponse(c *gin.Context) {
	ErrorResponse(c, http.StatusForbidden, "forbidden")
}

func NotFoundResponse(c *gin.Context) {
	ErrorResponse(c, http.StatusNotFound, "resource not found")
}

func InternalErrorResponse(c *gin.Context) {
	ErrorResponse(c, http.StatusInternalServerError, "internal server error")
}
```

---

## Main Application

### `cmd/server/main.go`

```go
package main

import (
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"yourmodule/internal/api"
	"yourmodule/internal/config"
	"yourmodule/internal/db"
	"yourmodule/internal/middleware"
)

func main() {
	// Load configuration
	cfg := config.Load()

	// Initialize database
	if err := db.Init(cfg.Database.DSN()); err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}
	defer db.Close()

	// Create Gin router
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	// CORS middleware
	router.Use(cors.New(cors.Config{
		AllowOrigins:     cfg.CORS.AllowedOrigins,
		AllowMethods:     cfg.CORS.AllowedMethods,
		AllowHeaders:     cfg.CORS.AllowedHeaders,
		AllowCredentials: true,
	}))

	// Health check
	router.GET("/health", func(c *gin.Context) {
		api.OKResponse(c, gin.H{"status": "healthy"}, "")
	})

	// Setup routes
	api.SetupRoutes(router)

	// Start server
	log.Printf("Starting server on :%s", cfg.Server.Port)
	if err := router.Run(":" + cfg.Server.Port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
```

### `internal/api/routes.go`

```go
package api

import (
	"github.com/gin-gonic/gin"
	"yourmodule/internal/api/handlers"
	"yourmodule/internal/middleware"
)

func SetupRoutes(router *gin.Engine) {
	// Public routes
	public := router.Group("/api/v1")
	{
		// Authentication
		public.POST("/auth/register", handlers.Register)
		public.POST("/auth/login", handlers.Login)
		public.POST("/auth/refresh", handlers.RefreshToken)

		// Content (public access)
		public.GET("/kana", handlers.GetAllKana)
		public.GET("/kanji", handlers.GetAllKanji)
		public.GET("/vocabulary", handlers.GetAllVocabulary)
	}

	// Protected routes
	protected := router.Group("/api/v1")
	protected.Use(middleware.AuthMiddleware())
	{
		// User routes
		protected.GET("/auth/me", handlers.GetCurrentUser)
		protected.POST("/auth/logout", handlers.Logout)

		// Progress and stats
		protected.GET("/stats", handlers.GetStats)
		protected.POST("/progress/update", handlers.UpdateProgress)
		protected.POST("/progress/answer", handlers.RecordAnswer)
		protected.GET("/progress/history", handlers.GetProgressHistory)

		// Preferences
		protected.GET("/preferences", handlers.GetPreferences)
		protected.PUT("/preferences", handlers.UpdatePreferences)

		// Achievements
		protected.GET("/achievements", handlers.GetAchievements)
		protected.GET("/achievements/unlocked", handlers.GetUnlockedAchievements)
	}
}
```

---

## Running the Server

### `Makefile`

```makefile
.PHONY: help dev build test migrate-up migrate-down

help:
	@echo "Available commands:"
	@echo "  make dev              - Run development server with hot reload"
	@echo "  make build            - Build the application"
	@echo "  make test             - Run tests"
	@echo "  make migrate-up       - Run database migrations"
	@echo "  make migrate-down     - Rollback database migrations"

dev:
	air

build:
	go build -o bin/server cmd/server/main.go

test:
	go test ./...

migrate-up:
	migrate -path migrations -database "postgres://user:password@localhost/kanadojo?sslmode=disable" up

migrate-down:
	migrate -path migrations -database "postgres://user:password@localhost/kanadojo?sslmode=disable" down
```

### Running

```bash
# Development with hot reload
make dev

# Or manually
go run cmd/server/main.go

# Production build
make build
./bin/server
```

---

## Next Steps

1. **Set up handlers** in `internal/api/handlers/` for each feature
2. **Create services** in `internal/services/` with business logic
3. **Create repositories** in `internal/repository/` for database access
4. **Write unit tests** for services and handlers
5. **Set up database migrations**
6. **Test API endpoints** with Postman

This foundation provides a solid, scalable Go backend!
