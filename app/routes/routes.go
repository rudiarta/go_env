package routes

import (
	"github.com/rudirahardian/go_env/app/controller"
	"github.com/gin-gonic/gin"
)

func RouteInit(port string){
	// gin.SetMode(gin.ReleaseMode)
	gin.SetMode(gin.DebugMode)
	router := gin.Default()

	v1User := router.Group("/api/v1/user")
	v1User.GET("/wait", controller.V1UserWait)
	v1User.GET("/new", controller.V1UserNew)
	
	v2User := router.Group("/api/v2/user")
	v2User.GET("/wait", controller.V2UserWait)
	v2User.GET("/new", controller.V2UserNew)
	
	router.Run(":"+port)
}