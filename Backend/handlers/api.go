package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/fachru/backend/database"
)

type APIHandler struct {
	db *database.PostgresDB
}

func NewAPIHandler(db *database.PostgresDB) *APIHandler {
	return &APIHandler{db: db}
}

func (h *APIHandler) HealthCheck(c *gin.Context) {
	if err := h.db.GetDB().Ping(); err != nil {
		c.JSON(http.StatusServiceUnavailable, gin.H{
			"status": "unhealthy",
			"error":  "Database connection failed",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  "healthy",
		"message": "Database connection is working",
	})
}

func (h *APIHandler) GetLogs(c *gin.Context) {
	limitStr := c.DefaultQuery("limit", "10")
	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		limit = 10
	}

	logs, err := h.db.GetRecentLogs(limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to retrieve logs",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"logs": logs,
	})
}

func (h *APIHandler) CreateLog(c *gin.Context) {
	var req struct {
		Message string `json:"message" binding:"required"`
		Level   string `json:"level"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request body",
		})
		return
	}

	if req.Level == "" {
		req.Level = "info"
	}

	if err := h.db.LogMessage(req.Message, req.Level); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to log message",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Log entry created successfully",
	})
}

func (h *APIHandler) GetDashboardData(c *gin.Context) {
	// Implementasi untuk mengambil data dashboard
	c.JSON(http.StatusOK, gin.H{
		"data": "dashboard data here",
	})
}

func (h *APIHandler) GetMapData(c *gin.Context) {
	// Implementasi untuk mengambil data map
	c.JSON(http.StatusOK, gin.H{
		"data": "map data here",
	})
}