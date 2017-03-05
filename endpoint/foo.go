package endpoint

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/javinc/graham/domain"
	"github.com/javinc/graham/model"
)

// Foo handler
func Foo(c *gin.Context) {
	switch c.Request.Method {
	case http.MethodGet:
		switch c.Param("id") {
		case "":
			FindFoo(c)
		default:
			GetFoo(c)
		}
	case http.MethodPost:
		CreateFoo(c)
	case http.MethodPatch:
		UpdateFoo(c)
	case http.MethodDelete:
		RemoveFoo(c)
	}
}

// FindFoo endpoint handler
func FindFoo(c *gin.Context) {
	o, err := domain.FindFoo(c)
	if err != nil {
		c.JSON(400, &model.Error{
			Message: err.Error(),
		})

		return
	}

	c.JSON(200, o)
}

// GetFoo endpoint handler
func GetFoo(c *gin.Context) {
	id := c.Param("id")
	o, err := domain.GetFoo(c, id)
	if err != nil {
		c.JSON(400, &model.Error{
			Message: err.Error(),
		})

		return
	}

	c.JSON(200, o)
}

// CreateFoo endpoint handler
func CreateFoo(c *gin.Context) {
	p := new(model.Foo)
	c.BindJSON(p)
	o, err := domain.CreateFoo(c, p)
	if err != nil {
		c.JSON(400, &model.Error{
			Message: err.Error(),
		})

		return
	}

	c.JSON(200, o)
}

// UpdateFoo endpoint handler
func UpdateFoo(c *gin.Context) {
	p := new(model.Foo)
	c.BindJSON(p)
	p.ID = c.Param("id")
	o, err := domain.UpdateFoo(c, p)
	if err != nil {
		c.JSON(400, &model.Error{
			Message: err.Error(),
		})

		return
	}

	c.JSON(200, o)
}

// RemoveFoo endpoint handler
func RemoveFoo(c *gin.Context) {
	id := c.Param("id")
	o, err := domain.RemoveFoo(c, id)
	if err != nil {
		c.JSON(400, &model.Error{
			Message: err.Error(),
		})

		return
	}

	c.JSON(200, o)
}
