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
	r.POST("/register", endpoint.Register)
	r.POST("/login", endpoint.Login)

	// private
	p := r.Group("/", middleware.Auth(checkUser))
	{
		p.GET("/me", endpoint.Me)

		p.GET("/foo", endpoint.FindFoo)
		p.POST("/foo", endpoint.CreateFoo)
		p.GET("/foo/:id", endpoint.GetFoo)
		p.PATCH("/foo/:id", endpoint.UpdateFoo)
		p.DELETE("/foo/:id", endpoint.RemoveFoo)
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
		ID: payload["id"].(string),
	}

	// check if user id exists
	u, err := domain.GetUser(c, u.ID)
	if err != nil {
		return errors.New("user has no access to this endpoint")
	}

	// overrides initial instance on baseMiddleware in favor of user info
	initContext(c, u)

	return nil
}

func initContext(c *gin.Context, u *model.User) {
	// domain context creation
	domain.ToContext(c, logic.New(c, u))
}
