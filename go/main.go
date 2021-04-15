package main

import (
	controller "github.com/NUTFes/shift-app/controller"
	"github.com/NUTFes/shift-app/model"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	db := model.SetupModels()

	r.Use(func(c *r.Context)) {
		c.Set("db", db)
		c.Next()
	}

	r.GET("/books", controller.FindBooks)
	r.POST("/books", controller.CreateBook)
	r.GET("/books/:id", controller.FindBook)
	r.PATCH("/books/:id", controller.UpdateBook)
	r.DELETE("/books/:id", controller.DeleteBook)
	r.Run(":3000")
}
