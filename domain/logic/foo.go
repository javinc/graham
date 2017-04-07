package logic

import (
	"strings"

	"github.com/javinc/graham/model"
	"github.com/javinc/mango/errors"
)

func (x *logic) FindFoo(o *model.FooOpts) ([]*model.Foo, error) {
	r, err := x.Data.FindFoo(o)
	if err != nil {
		return r, err
	}

	return r, nil
}

func (x *logic) FindOneFoo(o *model.FooOpts) (*model.Foo, error) {
	d := new(model.Foo)

	// get only one record
	o.Slice = "0,1"
	r, err := x.FindFoo(o)
	if err != nil {
		return d, err
	}

	if len(r) == 0 {
		return d, errors.New("LOGIC_FOO_FIND1", "record not found")
	}

	return r[0], nil
}

func (x *logic) GetFoo(id string) (*model.Foo, error) {
	// validation
	if id == "" {
		return new(model.Foo), errors.
			New("LOGIC_FOO_GET_CHK", "id param is required")
	}

	r, err := x.Data.GetFoo(id)
	if err != nil {
		return r, err
	}

	return r, nil
}

func (x *logic) CreateFoo(p *model.Foo) (*model.Foo, error) {
	// validation
	if p.Title == "" {
		return p, errors.New("LOGIC_FOO_CREATE_CHK", "title field is required")
	}

	// modification
	p.Title = strings.ToUpper(p.Title)

	// write
	r, err := x.Data.CreateFoo(p)
	if err != nil {
		return r, err
	}

	return r, nil
}

func (x *logic) UpdateFoo(p *model.Foo) (*model.Foo, error) {
	// validation
	if p.ID == "" {
		return p, errors.New("LOGIC_FOO_UPDATE_CHK", "id field is required")
	}

	// write
	p, err := x.Data.UpdateFoo(p)
	if err != nil {
		return p, err
	}

	return p, nil
}

func (x *logic) RemoveFoo(id string) (*model.Foo, error) {
	// validation
	if id == "" {
		return new(model.Foo), errors.
			New("LOGIC_FOO_REMOVE_CHK", "id param is required")
	}

	// write
	r, err := x.Data.RemoveFoo(id)
	if err != nil {
		return r, err
	}

	return r, nil
}
