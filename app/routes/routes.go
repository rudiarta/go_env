package routes

import (
	"github.com/rudirahardian/go_env/app/controller"
	"github.com/gin-gonic/gin"
)

func RouteInit(){
	// gin.SetMode(gin.ReleaseMode)
	gin.SetMode(gin.DebugMode)
	router := gin.Default()

	v1User := router.Group("/api/v1/user")
	v1User.GET("/wait", controller.UserWait)
	v1User.GET("/new", controller.UserNew)
	
	router.Run()
}