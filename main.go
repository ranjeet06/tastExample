package main

import (
	"github.com/gin-gonic/gin"
	"testProject/controllers"
)

func main() {
	r := gin.Default()

	r.GET("/books", controllers.FindBooks)
	r.GET("/books/:id", controllers.FindBook)
	r.POST("/books", controllers.CreateBook)
	r.PATCH("/books/:id", controllers.UpdateBook)
	r.DELETE("/books/:id", controllers.DeleteBook)
	
	err := r.Run("localhost:8080")
	if err != nil {
		panic("not run")
	}

}
