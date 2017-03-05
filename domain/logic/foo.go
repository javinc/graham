package logic

import (
	"errors"
	"strings"

	"github.com/javinc/graham/model"
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
		return p, errors.New("title is required")
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
		return p, errors.New("id is required")
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
