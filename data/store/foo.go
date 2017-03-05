package store

import (
	"github.com/javinc/graham/model"
)

const tableName = "foo"

func (x *store) FindFoo() ([]*model.Foo, error) {
	l := []*model.Foo{}
	err := Find(tableName, &l)

	return l, err
}

func (x *store) FindOneFoo() ([]*model.Foo, error) {
	l := []*model.Foo{}
	err := FindOne(tableName, &l)

	return l, err
}

func (x *store) GetFoo(id string) (*model.Foo, error) {
	r := new(model.Foo)
	err := Get(tableName, id, &r)

	return r, err
}

func (x *store) CreateFoo(r *model.Foo) (*model.Foo, error) {
	id, err := Create(tableName, r)
	if err != nil {
		return r, err
	}

	r.ID = id

	return r, nil
}
