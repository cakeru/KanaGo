package api

import (
	"github.com/cakeru/kanago-backend/internal/api/handlers"
	"github.com/cakeru/kanago-backend/internal/api/middleware"
	"github.com/cakeru/kanago-backend/internal/config"
	"github.com/gin-gonic/gin"
)

// SetupRoutes configures all API routes
func SetupRoutes(router *gin.Engine, cfg *config.Config) {
	// Create auth handler
	authHandler := handlers.NewAuthHandler(cfg)

	// Public routes (no authentication required)
	public := router.Group("/api")
	{
		public.POST("/register", authHandler.Register)
		public.POST("/login", authHandler.Login)
		public.POST("/refresh", authHandler.RefreshToken)

		// --- kana routers ---
		kanaHandler := handlers.NewKanaHandler()
		public.GET("/kana", kanaHandler.GetAllKana)
		public.GET("/kana/:id", kanaHandler.GetKana)
		public.GET("/kana/type/:type", kanaHandler.GetKanaByType)
	}

	protected := router.Group("/api")
	protected.Use(middleware.AuthMiddleware(cfg))
	{
		// Protected routes (authentication required)
	}
}