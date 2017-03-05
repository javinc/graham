package store

import (
	"github.com/imdario/mergo"
	"github.com/javinc/graham/model"
)

const name = "foo"

func (x *store) FindFoo() ([]*model.Foo, error) {
	r := []*model.Foo{}
	err := Find(name, &l)

	return r, err
}

func (x *store) FindOneFoo() ([]*model.Foo, error) {
	r := []*model.Foo{}
	err := FindOne(name, &l)

	return r, err
}

func (x *store) GetFoo(id string) (*model.Foo, error) {
	r := new(model.Foo)
	err := Get(name, id, &r)

	return r, err
}

func (x *store) CreateFoo(p *model.Foo) (*model.Foo, error) {
	// default

	// meta

	id, err := Create(name, p)
	if err != nil {
		return p, err
	}

	p.ID = id

	return p, nil
}

func (x *store) UpdateFoo(p *model.Foo) (*model.Foo, error) {
	r, err := x.GetFoo(p.ID)
	if err != nil {
		return r, err
	}

	// meta

	id := p.ID
	p.ID = ""
	err = Update(name, id, p)
	if err != nil {
		return r, err
	}

	// merge old values with the new
	mergo.MergeWithOverwrite(r, p)

	return r, nil
}

func (x *store) RemoveFoo(id string) (*model.Foo, error) {
	r, err := x.GetFoo(id)
	if err != nil {
		return r, err
	}

	err = Remove(name, id)
	if err != nil {
		return r, err
	}

	return r, nil
}
