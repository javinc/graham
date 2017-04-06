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
	type form struct {
		Email string `json:"email"`
		Pass  string `json:"password"`
	}

	p := new(form)
	err := util.ParsePayload(c, &p)
	if err != nil {
		util.OutputError(c, err)

		return
	}

	o, err := domain.LoginUser(c, p.Email, p.Pass)
	util.Output(c, o, err)
}

// Me handler
func Me(c *gin.Context) {
	o, err := domain.CurrentUser(c)
	if err != nil {
		util.OutputError(c, err)

		return
	}

	util.Output(c, o, err)
}
