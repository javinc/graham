package main

import (
	"github.com/gin-gonic/gin"
	"github.com/javinc/mango/config"
	"github.com/javinc/mango/server"

	"github.com/javinc/graham/data"
	"github.com/javinc/graham/data/store"
	"github.com/javinc/graham/domain"
	"github.com/javinc/graham/domain/logic"
	"github.com/javinc/graham/endpoint/router"
	"github.com/javinc/graham/model"
)

func main() {
	r := server.Engine()

	// middlewares
	r.Use(middleware())
	r.Use(gin.Logger())

	// routes
	router.Routes(r)

	r.Run(config.GetString("host"))
}

func middleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		u := getUser()

		// data context creation
		data.ToContext(c, store.New(u))

		// domain context creation
		domain.ToContext(c, logic.New(
			u,
			data.FromContext(c),
		))

		c.Next()
	}
}

func getUser() *model.User {
	u := new(model.User)
	u.ID = "100"
	u.Type = "client"
	u.Email = "test@koko.com"

	return u
}
