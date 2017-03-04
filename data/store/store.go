package store

import (
	"github.com/javinc/mango/data"
	"github.com/javinc/mango/model"
)

type store struct {
	User *model.User
}

// New create new store object
func New(u *model.User) data.Data {
	d := &store{u}

	return d
}
