package store

import (
	"github.com/javinc/graham/model"
)

const tableName = "foo"

func (x *store) FindFoo() ([]*model.Foo, error) {
	l := []*model.Foo{}

	Find(tableName, &l)

	return l, nil
}

func (x *store) FindOneFoo() ([]*model.Foo, error) {
	l := []*model.Foo{}

	FindOne(tableName, &l)

	return l, nil
}

func (x *store) GetFoo(id string) (*model.Foo, error) {
	r := new(model.Foo)

	Get(tableName, id, &r)

	return r, nil
}

func (x *store) CreateFoo(r *model.Foo) (*model.Foo, error) {
	// assuming its inserted to database
	r.ID = "200"

	return r, nil
}
