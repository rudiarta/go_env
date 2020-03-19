package controller

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/rudirahardian/go_env/app/service"
)

func V1UserWait(c *gin.Context) {
	var user service.User
	user.Name = "asd"
	user.Age = 12
	var hasil = service.Ambil(user)
	c.JSON(http.StatusCreated, gin.H{"message": "V1 Todo item created successfully!", "rudi": hasil.GetUser()})
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