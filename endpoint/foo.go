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

// UpdateFoo endpoint handler
func UpdateFoo(c *gin.Context) {
	testID := "f0fafea2-88b5-4471-9d86-79c5196149d4"
	testPayload := &model.Foo{
		ID:          testID,
		Title:       "modified foo",
		Description: "MODIDIFIED foo",
		Age:         100,
		Taken:       model.TruePtr,
	}

	d, err := domain.UpdateFoo(c, testPayload)
	if err != nil {
		c.JSON(400, &model.Error{
			Message: err.Error(),
		})

		return
	}

	c.JSON(200, d)
}
