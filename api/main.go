package main

import (
	controllers "github.com/NUTFes/shift-app/controllers"
	"github.com/NUTFes/shift-app/models"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	db := models.SetupGorm()

	r.Use(func(c *gin.Context){
		c.Set("db", db)
		c.Next()
	})

	r.GET("/books", controllers.FindBooks)
	r.POST("/books", controllers.CreateBook)
	r.GET("/books/:id", controllers.FindBook)
	r.PATCH("/books/:id", controllers.UpdateBook)
	r.DELETE("/books/:id", controllers.DeleteBook)
	r.Run(":3000")
}
