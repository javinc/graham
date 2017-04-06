package store

import (
	"github.com/javinc/graham/data/rethink"
	"github.com/javinc/graham/model"
)

const (
	userTableName = "user"
)

func init() {
	rethink.CreateTable(userTableName)
}

func (x *store) FindUser() ([]*model.User, error) {
	l := []*model.User{}

	r := new(model.User)
	r.ID = "testUser"
	r.Email = "list Email"

	return l, nil
}

func (x *store) GetUser(id string) (*model.User, error) {
	r := new(model.User)
	r.ID = "testUser"
	r.Email = "detail Email"

	return r, nil
}
