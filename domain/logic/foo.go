package logic

import (
	"strings"

	"github.com/javinc/graham/model"
)

const (
	fooErrFind        = "DOMAIN_FOO_FIND"
	fooErrGet         = "DOMAIN_FOO_GET"
	fooErrCreate      = "DOMAIN_FOO_CREATE"
	fooErrCreateCheck = "DOMAIN_FOO_CREATE_CHECK"
	fooErrUpdate      = "DOMAIN_FOO_UPDATE"
	fooErrUpdateCheck = "DOMAIN_FOO_UPDATE_CHECK"
	fooErrRemove      = "DOMAIN_FOO_REMOVE"
	fooErrRemoveCheck = "DOMAIN_FOO_REMOVE_CHECK"
)

func (x *logic) FindFoo() ([]*model.Foo, error) {
	r, err := x.Data.FindFoo()
	if err != nil {
		return r, err
	}

	return r, nil
}

func (x *logic) GetFoo(id string) (*model.Foo, error) {
	r, err := x.Data.GetFoo(id)
	if err != nil {
		return r, err
	}

	return r, nil
}

func (x *logic) CreateFoo(p *model.Foo) (*model.Foo, error) {
	// validation
	if p.Title == "" {
		return p, &model.Error{
			Code:    fooErrCreateCheck,
			Message: "title field is required",
		}
	}

	// modification
	p.Title = strings.ToUpper(p.Title)

	// write
	r, err := x.Data.CreateFoo(p)
	if err != nil {
		return r, &model.Error{
			Code:    fooErrCreate,
			Message: err.Error(),
		}
	}

	return r, nil
}

func (x *logic) UpdateFoo(p *model.Foo) (*model.Foo, error) {
	// validation
	if p.ID == "" {
		return p, &model.Error{
			Code:    fooErrUpdateCheck,
			Message: "id field is required",
		}
	}

	// write
	p, err := x.Data.UpdateFoo(p)
	if err != nil {
		return p, err
	}

	return p, nil
}

func (x *logic) RemoveFoo(id string) (*model.Foo, error) {
	// write
	r, err := x.Data.RemoveFoo(id)
	if err != nil {
		return r, err
	}

	return r, nil
}
