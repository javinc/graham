package main

import (
	"github.com/gin-gonic/gin"
	"github.com/javinc/graham/data"
	"github.com/javinc/graham/data/store"
	"github.com/javinc/graham/domain"
	"github.com/javinc/graham/domain/logic"
	"github.com/javinc/graham/endpoint"
	"github.com/javinc/graham/model"
)

func main() {
	r := gin.Default()

	// middlewares
	r.Use(middleware())

	// routes
	routes(r)

	r.Run(":8000")
}

func routes(r *gin.Engine) {
	r.Any("/foo", endpoint.Foo)
	r.Any("/foo/:id", endpoint.Foo)
	// r.GET("/foo", endpoint.FindFoo)
	// r.POST("/foo", endpoint.CreateFoo)
	// r.GET("/foo/:id", endpoint.GetFoo)
	// r.PATCH("/foo/:id", endpoint.UpdateFoo)
	// r.DELETE("/foo/:id", endpoint.RemoveFoo)

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
