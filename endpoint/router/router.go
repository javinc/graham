package router

import (
	"github.com/gin-gonic/gin"

	"github.com/javinc/graham/endpoint"
)

// Routes endpoint routes
func Routes(r *gin.Engine) {
	r.GET("/foo", endpoint.FindFoo)
	r.POST("/foo", endpoint.CreateFoo)
	r.GET("/foo/:id", endpoint.GetFoo)
	r.PATCH("/foo/:id", endpoint.UpdateFoo)
	r.DELETE("/foo/:id", endpoint.RemoveFoo)

	r.GET("/user", endpoint.FindUser)
	r.GET("/user/:id", endpoint.GetUser)
}
