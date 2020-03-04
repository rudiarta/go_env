package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()

	v1 := router.Group("/api/v1/todos")
	{
		v1.GET("/", createTodo)
	}
	router.Run()

}

// createTodo add a new todo
func createTodo(c *gin.Context) {
	c.JSON(http.StatusCreated, gin.H{"message": "Todo item created successfully!", "resourceId": 1})
}