package logic

import (
	"log"

	"github.com/gin-gonic/gin"

	"github.com/javinc/graham/data"
	"github.com/javinc/graham/data/store"
	"github.com/javinc/graham/domain"
	"github.com/javinc/graham/model"
)

type logic struct {
	// user accessor
	User *model.User

	// data store context
	Data data.Data
}

// New create new domain object
func New(c *gin.Context, u *model.User) domain.Domain {
	// data context creation
	data.ToContext(c, store.New(u))
	log.Println(data.FromContext(c))

	d := &logic{
		u,
		data.FromContext(c),
	}

	return d
}
