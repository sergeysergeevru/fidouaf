package main

import (
	"github.com/sergeysergeevru/fidouaf/pkg/router"
)


func main() {
	r := router.SetupRouter()
	r.Static("/static","./public/")
	r.Run(":80") // listen and serve on 0.0.0.0:8080
}
