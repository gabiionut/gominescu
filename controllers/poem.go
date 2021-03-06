package controllers

import (
	"fmt"
	"net/http"

	"github.com/gabiionut/gominescu/models"
	"github.com/gin-gonic/gin"
)

// GetPoems Get all poems
func GetPoems(c *gin.Context) {
	var poems []models.Poem
	models.DB.Find(&poems)

	c.JSON(http.StatusOK, gin.H{"data": poems, "count": len(poems)})
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
	var poems []models.Poem
	key := c.Query("key")

	if err := models.DB.Where("title LIKE ?", fmt.Sprintf("%%%s%%", key)).Find(&poems).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": poems, "count": len(poems)})
}

func AddPoem(c *gin.Context) {
	var input models.Poem

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Add poem

	poem := models.Poem{Title: input.Title, Content: input.Content}
	models.DB.Create(&poem)

	c.JSON(http.StatusOK, gin.H{"data": poem})
}
