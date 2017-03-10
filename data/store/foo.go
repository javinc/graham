package store

import (
	"github.com/imdario/mergo"
	"github.com/javinc/graham/data/rethink"
	"github.com/javinc/graham/model"
)

const (
	fooTableName = "foo"

	fooErrFind        = "DATA_FOO_FIND"
	fooErrFindOne     = "DATA_FOO_FIND_ONE"
	fooErrGet         = "DATA_FOO_GET"
	fooErrCreate      = "DATA_FOO_CREATE"
	fooErrUpdate      = "DATA_FOO_UPDATE"
	fooErrUpdateCheck = "DATA_FOO_UPDATE_CHECK"
	fooErrRemove      = "DATA_FOO_REMOVE"
	fooErrRemoveCheck = "DATA_FOO_REMOVE_CHECK"
)

func init() {
	rethink.CreateTable(fooTableName)
}

func (x *store) FindFoo() ([]*model.Foo, error) {
	r := []*model.Foo{}
	err := rethink.Find(fooTableName, &r)
	if err != nil {
		return r, &model.Error{
			Code:    fooErrFind,
			Message: err.Error(),
		}
	}

	return r, nil
}

func (x *store) FindOneFoo() ([]*model.Foo, error) {
	r := []*model.Foo{}
	err := rethink.FindOne(fooTableName, &r)
	if err != nil {
		return r, &model.Error{
			Code:    fooErrFindOne,
			Message: err.Error(),
		}
	}

	return r, nil
}

func (x *store) GetFoo(id string) (*model.Foo, error) {
	r := new(model.Foo)
	err := rethink.Get(fooTableName, id, &r)
	if err != nil {
		return r, &model.Error{
			Code:    fooErrGet,
			Message: err.Error(),
		}
	}

	return r, nil
}

func (x *store) CreateFoo(p *model.Foo) (*model.Foo, error) {
	// default

	// meta

	id, err := rethink.Create(fooTableName, p)
	if err != nil {
		return p, &model.Error{
			Code:    fooErrCreate,
			Message: err.Error(),
		}
	}

	p.ID = id

	return p, nil
}

func (x *store) UpdateFoo(p *model.Foo) (*model.Foo, error) {
	r, _ := x.GetFoo(p.ID)
	if r.ID == "" {
		return r, &model.Error{
			Code:    fooErrUpdateCheck,
			Message: "record does not exist",
		}
	}

	// meta

	id := p.ID
	p.ID = ""
	err := rethink.Update(fooTableName, id, p)
	if err != nil {
		return r, &model.Error{
			Code:    fooErrUpdate,
			Message: err.Error(),
		}
	}

	// merge old values with the new
	mergo.MergeWithOverwrite(r, p)

	return r, nil
}

func (x *store) RemoveFoo(id string) (*model.Foo, error) {
	r, _ := x.GetFoo(id)
	if r.ID == "" {
		return r, &model.Error{
			Code:    fooErrRemoveCheck,
			Message: "record does not exist",
		}
	}

	err := rethink.Remove(fooTableName, id)
	if err != nil {
		return r, &model.Error{
			Code:    fooErrRemove,
			Message: err.Error(),
		}
	}

	return r, nil
}
