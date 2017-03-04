package main

import (
	"github.com/gin-gonic/gin"
	"github.com/javinc/mango/data"
	"github.com/javinc/mango/data/store"
	"github.com/javinc/mango/domain"
	"github.com/javinc/mango/domain/logic"
	"github.com/javinc/mango/endpoint"
	"github.com/javinc/mango/model"
)

func main() {
	r := gin.Default()

	// middlewares
	r.Use(middleware())

	// routes
	routes(r)

	r.Run()
}

func routes(r *gin.Engine) {
	r.GET("/foo", endpoint.FindFoo)
	r.GET("/foo/:id", endpoint.GetFoo)
	r.POST("/foo", endpoint.CreateFoo)

	r.GET("/user", endpoint.FindUser)
	r.GET("/user/:id", endpoint.GetUser)
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
