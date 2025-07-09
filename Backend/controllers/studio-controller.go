package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/VeryFach/MMD-FILKOM-TINGAL/Backend/models"
)

// Dummy data sebagai contoh
var regions = []models.Region{}

// Get all regions
func GetRegions(c *gin.Context) {
	c.JSON(http.StatusOK, regions)
}

// Get region by ID
func GetRegionByID(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseInt(idParam, 10, 64)
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
	var newRegion models.Region
	if err := c.ShouldBindJSON(&newRegion); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	newRegion.ID = int64(len(regions) + 1)
	newRegion.CalculateTotalPopulation()
	regions = append(regions, newRegion)
	c.JSON(http.StatusCreated, newRegion)
}

// Update region
func UpdateRegion(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	var updatedRegion models.Region
	if err := c.ShouldBindJSON(&updatedRegion); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	for i, region := range regions {
		if region.ID == id {
			updatedRegion.ID = id
			updatedRegion.CalculateTotalPopulation()
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
	id, err := strconv.ParseInt(idParam, 10, 64)
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