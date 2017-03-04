package logic

import (
	"github.com/javinc/mango/data"
	"github.com/javinc/mango/domain"
	"github.com/javinc/mango/model"
)

type logic struct {
	// user accessor
	User *model.User

	// data store context
	Data data.Data
}

// New create new domain object
func New(u *model.User, store data.Data) domain.Domain {
	d := &logic{
		u,
		store,
	}

	return d
}
