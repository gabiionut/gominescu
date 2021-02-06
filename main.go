package main

import (
	"github.com/gabiionut/gominescu/controllers"
	"github.com/gabiionut/gominescu/models"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	models.ConnectDataBase()

	r.GET("/poems", controllers.GetPoems)
	r.GET("/poems/:id", controllers.GetPoemByID)
	r.GET("/search", controllers.SearchPoem)

	r.Run()
}
