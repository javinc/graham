package logic

import (
	"github.com/javinc/graham/data"
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
func New(u *model.User, store data.Data) domain.Domain {
	d := &logic{
		u,
		store,
	}

	return d
}
