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

			break
		}

		o, err = domain.GetFoo(c, id)
	case http.MethodPost:
		o, err = domain.CreateFoo(c, p)
	case http.MethodPatch:
		p.ID = id
		o, err = domain.UpdateFoo(c, p)
	case http.MethodDelete:
		o, err = domain.RemoveFoo(c, id)
	}

	output(c, o, err)
}
