package store

import (
	"github.com/javinc/mango/model"
)

func (x *store) FindFoo() ([]*model.Foo, error) {
	l := []*model.Foo{}

	r := new(model.Foo)
	r.Title = "list Title"
	r.Description = "list Description"

	l = append(l, r)
	l = append(l, r)
	l = append(l, r)

	return l, nil
}

func (x *store) GetFoo(id string) (*model.Foo, error) {
	r := new(model.Foo)
	r.Title = "detail title"
	r.Description = "detail Description"

	return r, nil
}

func (x *store) CreateFoo(r *model.Foo) (*model.Foo, error) {
	// assuming its inserted to database
	r.ID = "200"

	return r, nil
}
