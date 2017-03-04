package store

import (
	r "github.com/gorethink/gorethink"
	db "github.com/javinc/mango/database/rethink"

	"github.com/javinc/graham/model"
)

func (x *store) FindFoo() ([]*model.Foo, error) {
	l := []*model.Foo{}

	db.Find(r.Table("foo"), &l)

	return l, nil
}

func (x *store) GetFoo(id string) (*model.Foo, error) {
	r := new(model.Foo)
	r.Title = "detail title"
	r.Description = "detail description"

	return r, nil
}

func (x *store) CreateFoo(r *model.Foo) (*model.Foo, error) {
	// assuming its inserted to database
	r.ID = "200"

	return r, nil
}