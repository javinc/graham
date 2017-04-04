package router

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/javinc/mango/server/middleware"

	"github.com/javinc/graham/endpoint"
)

// Routes endpoint routes
func Routes(r *gin.Engine) {
	// public
	r.GET("/foo", endpoint.FindFoo)
	r.POST("/foo", endpoint.CreateFoo)
	r.GET("/foo/:id", endpoint.GetFoo)
	r.PATCH("/foo/:id", endpoint.UpdateFoo)
	r.DELETE("/foo/:id", endpoint.RemoveFoo)

	r.GET("/user", endpoint.FindUser)
	r.GET("/user/:id", endpoint.GetUser)

	// private
	p := r.Group("/", middleware.PrivateMiddleware(checkUser))
	{
		p.GET("/foox", endpoint.FindFoo)
	}
}

func checkUser(payload map[string]interface{}) error {
	id := payload["id"].(string)

	// check if id exists
	if id != "testUser" {
		return errors.New("THIS IS INVALID USER")
	}

	return nil
}
