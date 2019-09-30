package router

import (
	"github.com/gin-gonic/gin"
	"github.com/sergeysergeevru/fidouaf/pkg/uaf/shared"
)

func Registration(c *gin.Context) {
	var regRequest []shared.RegistrationRequest
	err := c.BindJSON(&regRequest)
	if err != nil {
		c.JSON(400, gin.H{
			"success": false,
			"message": err,
		})
		c.Abort()
		return
	}
	c.JSON(200, regRequest)
}
