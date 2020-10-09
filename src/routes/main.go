package routes

import (
	"batara/src/controlles"
	"batara/src/middlewares"
	"github.com/gin-gonic/gin"
)

func Routes(){
	router := gin.Default()
	v1 := router.Group("/v1")
	{
		v1.GET("/hallo", middlewares.TokenAuthMiddleware(),controlles.HitApi)
		v1.GET("/token", controlles.GetToken)
	}
	router.Run(":8000")
}
