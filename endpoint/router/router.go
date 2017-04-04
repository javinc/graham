package router

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/javinc/mango/server/middleware"

	"github.com/javinc/graham/endpoint"
	"github.com/javinc/graham/model"
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

	// catchers
	r.NoRoute(notFound)
}

func checkUser(payload map[string]interface{}) error {
	id := payload["id"].(string)

	// check if id exists
	if id != "testUser" {
		return errors.New("user has no access to this resource")
	}

	return nil
}

func notFound(c *gin.Context) {
	c.JSON(http.StatusNotFound, model.Error{
		Name:    "NOT_FOUND",
		Message: "resource not found",
	})
}
