package test

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

// createTodo add a new todo
func UserWait(c *gin.Context) {
	c.JSON(http.StatusCreated, gin.H{"message": "Todo item created successfully!", "rudi": 1})
}

func UserNew(c *gin.Context) {
	c.JSON(http.StatusCreated, gin.H{"message": "Todo item created successfully new!", "rudi": 1})
}