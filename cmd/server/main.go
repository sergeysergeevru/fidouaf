package main

import (
	"github.com/sergeysergeevru/fidouaf/internal/router"
)


func main() {
	r := router.SetupRouter()
	r.Static("/static","./public/")
	r.Run(":8480") // listen and serve on 0.0.0.0:8080
}
