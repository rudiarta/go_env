package controller

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func V1UserWait(c *gin.Context) {
	c.JSON(http.StatusCreated, gin.H{"message": "V1 Todo item created successfully!", "rudi": 1})
}

func V1UserNew(c *gin.Context) {
	c.JSON(http.StatusCreated, gin.H{"message": "V1 Todo item created successfully new!", "rudi": 1})
}

func V2UserWait(c *gin.Context) {
	c.JSON(http.StatusCreated, gin.H{"message": "V2 Todo item created successfully!", "rudi": 1})
}

func V2UserNew(c *gin.Context) {
	c.JSON(http.StatusCreated, gin.H{"message": "V2 Todo item created successfully new!", "rudi": 1})
}