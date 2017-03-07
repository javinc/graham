package endpoint

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/javinc/graham/domain"
	"github.com/javinc/graham/model"
)

// Foo handler
func Foo(c *gin.Context) {
	var err error
	var o interface{}
	var p *model.Foo

	id := c.Param("id")
	if c.Request.Method != http.MethodGet {
		p = new(model.Foo)
		c.BindJSON(p)
	}

	switch c.Request.Method {
	case http.MethodGet:
		if id == "" {
			o, err = domain.FindFoo(c)
			output(c, o, err)

			return
		}

		o, err = domain.GetFoo(c, id)
		output(c, o, err)
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
