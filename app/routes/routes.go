package routes

import (
	"github.com/rudirahardian/go_env/app/controller/test"
	"github.com/gin-gonic/gin"
)

func RouteInit(port string){
	// gin.SetMode(gin.ReleaseMode)
	gin.SetMode(gin.DebugMode)
	router := gin.Default()

	v1User := router.Group("/api/v1/user")
	v1User.GET("/wait", test.UserWait)
	v1User.GET("/new", test.UserNew)
	
	router.Run(":"+port)
}