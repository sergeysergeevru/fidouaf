package main

import (
	"github.com/sergeysergeevru/fidouaf/router"
)


func main() {
	r := router.SetupRouter()
	r.Run(":81") // listen and serve on 0.0.0.0:8080
}
