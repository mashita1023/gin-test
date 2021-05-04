package main

import (
  "fmt"
	"os"

  controllers "github.com/NUTFes/shift-app/controllers"
	"github.com/NUTFes/shift-app/models"
	"github.com/NUTFes/shift-app/middlewares"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

  mw := middlewares.SetupFirebase()
	db := models.SetupGorm()

	r.Use(func(c *gin.Context){
		c.Set("db", db)
		c.Next()
	})

  fmt.Printf("%v\n", true)
  fmt.Printf("%v\n", os.Getenv("FIRE_BASE_TYPE"))
	r.GET("/books", controllers.FindBooks)
	r.POST("/books", controllers.CreateBook)
	r.GET("/books/:id", controllers.FindBook)
	r.PATCH("/books/:id", controllers.UpdateBook)
	r.DELETE("/books/:id", controllers.DeleteBook)
	r.Run(":3000")
}
