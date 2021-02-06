package models

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// Poem Represents a poem
type Poem struct {
	ID      uint   `json:"id" binding:"required"`
	Title   string `json:"title" binding:"required"`
	Content string `json:"content" binding:"required"`
}

// DB Database context
var DB *gorm.DB

// ConnectDataBase Connect to database
func ConnectDataBase() {
	db, err := gorm.Open(sqlite.Open("gominescu.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	DB = db
}
