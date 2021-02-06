package controllers

import (
	"net/http"

	"github.com/gabiionut/gominescu/models"
	"github.com/gin-gonic/gin"
)

// GetPoems Get all poems
func GetPoems(c *gin.Context) {
	var poems []models.Poem
	models.DB.Find(&poems)

	c.JSON(http.StatusOK, gin.H{"data": poems})
}

// GetPoemByID Get poem by id
func GetPoemByID(c *gin.Context) {
	var poem models.Poem

	if err := models.DB.Where("id = ?", c.Param("id")).First(&poem).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": poem})
}

// SearchPoem Search poem
func SearchPoem(c *gin.Context) {
	var poem models.Poem
	key := c.Query("key")

	if err := models.DB.Where("title = ?", key).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": poem})
}
