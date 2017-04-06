package main

import (
	"github.com/gin-gonic/gin"
	"github.com/javinc/mango/config"
	"github.com/javinc/mango/server"

	"github.com/javinc/graham/endpoint/router"
)

func main() {
	m := gin.DebugMode

	// set engine mode to produciton
	if config.GetBool("live") {
		m = gin.ReleaseMode
	}

	r := server.Engine(m)

	// middlewares
	r.Use(gin.Logger())

	// routes
	router.Routes(r)

	r.Run(config.GetString("host"))
}
