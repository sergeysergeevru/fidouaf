package router

import "github.com/gin-gonic/gin"

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.POST("/uaf/reg", Registration)
	r.GET("/registration", WebRegistration)
	return r
}

