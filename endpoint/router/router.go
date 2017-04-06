package router

import (
	"errors"

	"github.com/gin-gonic/gin"

	"github.com/javinc/graham/domain"
	"github.com/javinc/graham/domain/logic"
	"github.com/javinc/graham/endpoint"
	"github.com/javinc/graham/endpoint/util"
	"github.com/javinc/graham/model"
	"github.com/javinc/mango/server/middleware"
)

// Routes endpoint routes
func Routes(r *gin.Engine) {
	r.Use(baseMiddleware())

	// public
	r.GET("/foo", endpoint.FindFoo)
	r.POST("/foo", endpoint.CreateFoo)
	r.GET("/foo/:id", endpoint.GetFoo)
	r.PATCH("/foo/:id", endpoint.UpdateFoo)
	r.DELETE("/foo/:id", endpoint.RemoveFoo)

	r.GET("/user", endpoint.FindUser)
	r.GET("/user/:id", endpoint.GetUser)

	// private
	p := r.Group("/", middleware.Auth(checkUser))
	{
		p.GET("/foox", endpoint.FindFoo)
	}

	// catcher
	r.NoRoute(util.OutputNotFound)
}

func baseMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		initContext(c, new(model.User))

		c.Next()
	}
}

func checkUser(c *gin.Context, payload map[string]interface{}) error {
	u := &model.User{
		ID:   payload["id"].(string),
		Name: payload["name"].(string),
	}

	// check if user id exists
	_, err := domain.GetUser(c, u.ID)
	if err != nil {
		return errors.New("user has no access to this endpoint")
	}

	initContext(c, u)

	return nil
}

func initContext(c *gin.Context, u *model.User) {
	// domain context creation
	domain.ToContext(c, logic.New(c, u))
}
