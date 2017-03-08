package endpoint

import (
	"github.com/gin-gonic/gin"
	"github.com/javinc/graham/domain"
	"github.com/javinc/graham/endpoint/util"
	"github.com/javinc/graham/model"
)

// FindFoo handler
func FindFoo(c *gin.Context) {
	o, err := domain.FindFoo(c)
	util.Output(c, o, err)
}

// FooGet handler
func GetFoo(c *gin.Context) {
	id := c.Param("id")
	o, err := domain.GetFoo(c, id)
	util.Output(c, o, err)
}

// CreateFoo handler
func CreateFoo(c *gin.Context) {
	p := new(model.Foo)
	err := util.ParsePayload(c, &p)
	if err != nil {
		util.OutputError(c, err)

		return
	}

	o, err := domain.CreateFoo(c, p)
	util.Output(c, o, err)
}

// UpdateFoo handler
func UpdateFoo(c *gin.Context) {
	p := new(model.Foo)
	err := util.ParsePayload(c, &p)
	if err != nil {
		util.OutputError(c, err)

		return
	}

	p.ID = c.Param("id")
	o, err := domain.UpdateFoo(c, p)
	util.Output(c, o, err)
}

// RemoveFoo handler
func RemoveFoo(c *gin.Context) {
	id := c.Param("id")
	o, err := domain.RemoveFoo(c, id)
	util.Output(c, o, err)
}
