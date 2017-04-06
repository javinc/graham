package endpoint

import (
	"github.com/gin-gonic/gin"

	"github.com/javinc/graham/domain"
	"github.com/javinc/graham/endpoint/util"
	"github.com/javinc/graham/model"
)

// Register handler
func Register(c *gin.Context) {
	p := new(model.User)
	err := util.ParsePayload(c, &p)
	if err != nil {
		util.OutputError(c, err)

		return
	}

	o, err := domain.CreateUser(c, p)
	util.Output(c, o, err)
}

// Login handler
func Login(c *gin.Context) {
	id := c.Param("id")
	o, err := domain.GetUser(c, id)
	util.Output(c, o, err)
}

// Me handler
func Me(c *gin.Context) {
	id := c.Param("id")
	o, err := domain.GetUser(c, id)
	util.Output(c, o, err)
}
