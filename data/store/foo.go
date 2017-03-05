package store

import (
	"github.com/imdario/mergo"
	"github.com/javinc/graham/model"
)

const name = "foo"

func (x *store) FindFoo() ([]*model.Foo, error) {
	l := []*model.Foo{}
	err := Find(name, &l)

	return l, err
}

func (x *store) FindOneFoo() ([]*model.Foo, error) {
	l := []*model.Foo{}
	err := FindOne(name, &l)

	return l, err
}

func (x *store) GetFoo(id string) (*model.Foo, error) {
	r := new(model.Foo)
	err := Get(name, id, &r)

	return r, err
}

func (x *store) CreateFoo(i *model.Foo) (*model.Foo, error) {
	id, err := Create(name, i)
	if err != nil {
		return i, err
	}

	// default

	i.ID = id

	return i, nil
}

func (x *store) UpdateFoo(in *model.Foo) (*model.Foo, error) {
	r, err := x.GetFoo(in.ID)
	if err != nil {
		return r, err
	}

	// meta

	id := in.ID
	in.ID = ""
	err = Update(name, id, in)
	if err != nil {
		return r, err
	}

	// merge old values with the new
	mergo.MergeWithOverwrite(&r, in)

	return in, nil
}
