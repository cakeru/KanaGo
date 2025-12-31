package handlers

import (
	"net/http"
	"strconv"

	"github.com/cakeru/kanago-backend/internal/services"
	"github.com/gin-gonic/gin"
)

// KanaHandler handles kana-related HTTP requests
type KanaHandler struct {
	kanaService *services.KanaService
}

// NewKanaHandler creates a new instace of kanahandler

func NewKanaHandler() *KanaHandler {
	return &KanaHandler{
		kanaService: services.NewKanaService(),
	}
}

// GetKana handles GET /api/kana/:id
func (h *KanaHandler) GetKana(c *gin.Context){
	// Get ID from URL parameter
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid kana ID",
		})
		return
	}

	// Call service
	resp, err := h.kanaService.GetKana(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	// Return success response
	c.JSON(http.StatusOK, resp)
}

// GetKanaByType handles GET /api/kana/type/:type
func (h *KanaHandler) GetKanaByType(c *gin.Context){
	// Get type from URL parameter
	kanaType := c.Param("type")

	// Call service
	resp, err := h.kanaService.GetKanaByType(kanaType)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// Return success response
	c.JSON(http.StatusOK, gin.H{
		"type": kanaType,
		"data": resp,
		"count": len(resp),
	})
}

// GetAllKana handles GET /api/kana
func (h *KanaHandler) GetAllKana(c *gin.Context){
	// Call service
	resp, err := h.kanaService.GetAllKana()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	// Return success response
	c.JSON(http.StatusOK, resp)
}