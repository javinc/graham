package endpoint

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/javinc/graham/domain"
	"github.com/javinc/graham/model"
)

// Foo handler
func Foo(c *gin.Context) {
	id := c.Param("id")
	p := new(model.Foo)
	c.BindJSON(p)

	var err error
	var o interface{}
	switch c.Request.Method {
	case http.MethodGet:
		switch id {
		case "":
			o, err = domain.FindFoo(c)
			output(c, o, err)

		default:
			o, err = domain.GetFoo(c, id)
			output(c, o, err)
		}

	case http.MethodPost:
		o, err = domain.CreateFoo(c, p)
		output(c, o, err)
	case http.MethodPatch:
		p.ID = id
		o, err = domain.UpdateFoo(c, p)
		output(c, o, err)
	case http.MethodDelete:
		o, err = domain.RemoveFoo(c, id)
		output(c, o, err)
	}
}
