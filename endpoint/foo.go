package endpoint

import (
	"github.com/gin-gonic/gin"
	"github.com/javinc/graham/domain"
	"github.com/javinc/graham/model"
)

// FindFoo endpoint handler
func FindFoo(c *gin.Context) {
	r, err := domain.FindFoo(c)
	if err != nil {
		c.JSON(400, &model.Error{
			Message: err.Error(),
		})

		return
	}

	c.JSON(200, r)
}

// GetFoo endpoint handler
func GetFoo(c *gin.Context) {
	id := c.Param("id")
	d, err := domain.GetFoo(c, id)
	if err != nil {
		c.JSON(400, &model.Error{
			Message: err.Error(),
		})

		return
	}

	c.JSON(200, d)
}

// CreateFoo endpoint handler
func CreateFoo(c *gin.Context) {
	testPayload := &model.Foo{
		Title:       "create foo",
		Description: "create foo desc",
		Age:         200,
	}

	d, err := domain.CreateFoo(c, testPayload)
	if err != nil {
		c.JSON(400, &model.Error{
			Message: err.Error(),
		})

		return
	}

	c.JSON(200, d)
}
