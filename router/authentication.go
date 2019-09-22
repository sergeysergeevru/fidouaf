package router

import (
	"github.com/gin-gonic/gin"
	"github.com/sergeysergeevru/fidouaf/webauthn"
)

func WebRegistration(c *gin.Context) {
	var resp webauthn.InitialRegistrationResponse
	c.JSON(200, resp)
}