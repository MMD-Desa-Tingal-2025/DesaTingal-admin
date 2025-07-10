package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/fachru/backend/models"
	"github.com/google/uuid"
)

// Dummy data sebagai contoh
var regions = []models.Area{}

// Get all regions
func GetRegions(c *gin.Context) {
	c.JSON(http.StatusOK, regions)
}

// Get region by ID
func GetRegionByID(c *gin.Context) {
	idParam := c.Param("id")
	id, err := uuid.Parse(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	for _, region := range regions {
		if region.ID == id {
			c.JSON(http.StatusOK, region)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "Region not found"})
}

// Create new region
func CreateRegion(c *gin.Context) {
	var newRegion models.Area
	if err := c.ShouldBindJSON(&newRegion); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newRegion.ID = uuid.New()
	regions = append(regions, newRegion)
	c.JSON(http.StatusCreated, newRegion)
}

// Update existing region
func UpdateRegion(c *gin.Context) {
	idParam := c.Param("id")
	id, err := uuid.Parse(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var updatedRegion models.Area
	if err := c.ShouldBindJSON(&updatedRegion); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	for i, region := range regions {
		if region.ID == id {
			updatedRegion.ID = id
			regions[i] = updatedRegion
			c.JSON(http.StatusOK, updatedRegion)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "Region not found"})
}

// Delete region
func DeleteRegion(c *gin.Context) {
	idParam := c.Param("id")
	id, err := uuid.Parse(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	for i, region := range regions {
		if region.ID == id {
			regions = append(regions[:i], regions[i+1:]...)
			c.JSON(http.StatusOK, gin.H{"message": "Region deleted"})
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "Region not found"})
}
