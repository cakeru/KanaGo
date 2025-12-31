package main

import (
	"fmt"
	"log"

	"github.com/cakeru/kanago-backend/internal/api"
	"github.com/cakeru/kanago-backend/internal/config"
	"github.com/cakeru/kanago-backend/internal/db"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}
}

func main() {

	// Load configuration
    cfg := config.Load()
	// Initialize database
	if err := db.Init(cfg); err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}
	defer db.Close()

	r := gin.Default()

	r.Use(cors.Default())

	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "Server is running",
		})
	})

	// Setup API routes
	api.SetupRoutes(r, cfg)

	// Start server
	port := ":" + cfg.Server.Port
	fmt.Println("Server is running on port " + port)
	if err := r.Run(port); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}