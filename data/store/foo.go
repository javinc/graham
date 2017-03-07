package store

import (
	"github.com/imdario/mergo"
	"github.com/javinc/graham/model"
)

const (
	fooTableName = "foo"

	fooFindErr    = "DATA_STORE_FIND_FOO"
	fooFindOneErr = "DATA_STORE_FIND_ONE_FOO"
)

func (x *store) FindFoo() ([]*model.Foo, error) {
	r := []*model.Foo{}
	err := Find(fooTableName, &r)
	if err != nil {
		return r, &model.Error{
			Code:    fooFindErr,
			Message: err.Error(),
		}
	}

	return r, nil
}

func (x *store) FindOneFoo() ([]*model.Foo, error) {
	r := []*model.Foo{}
	err := FindOne(fooTableName, &r)
	if err != nil {
		return r, &model.Error{
			Code:    fooFindOneErr,
			Message: err.Error(),
		}
	}

	return r, nil
}

func (x *store) GetFoo(id string) (*model.Foo, error) {
	r := new(model.Foo)
	err := Get(fooTableName, id, &r)
	if err != nil {
		return r, &model.Error{
			Code:    "DATA_STORE_GET_FOO",
			Message: err.Error(),
		}
	}

	return r, nil
}

func (x *store) CreateFoo(p *model.Foo) (*model.Foo, error) {
	// default

	// meta

	id, err := Create(fooTableName, p)
	if err != nil {
		return p, &model.Error{
			Code:    "DATA_STORE_CREATE_FOO",
			Message: err.Error(),
		}
	}

	p.ID = id

	return p, nil
}

func (x *store) UpdateFoo(p *model.Foo) (*model.Foo, error) {
	r, err := x.GetFoo(p.ID)
	if err != nil {
		return r, &model.Error{
			Code:    "DATA_STORE_UPDATE_FOO_CHECK",
			Message: err.Error(),
		}
	}

	// meta

	id := p.ID
	p.ID = ""
	err = Update(fooTableName, id, p)
	if err != nil {
		return r, &model.Error{
			Code:    "DATA_STORE_UPDATE_FOO",
			Message: err.Error(),
		}
	}

	// merge old values with the new
	mergo.MergeWithOverwrite(r, p)

	return r, nil
}

func (x *store) RemoveFoo(id string) (*model.Foo, error) {
	r, err := x.GetFoo(id)
	if err != nil {
		return r, &model.Error{
			Code:    "DATA_STORE_DELETE_FOO_CHECK",
			Message: err.Error(),
		}
	}

	err = Remove(fooTableName, id)
	if err != nil {
		return r, &model.Error{
			Code:    "DATA_STORE_DELETE_FOO",
			Message: err.Error(),
		}
	}

	return r, nil
}
