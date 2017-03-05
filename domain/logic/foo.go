package logic

import (
	"errors"
	"strings"

	"github.com/javinc/graham/model"
)

func (x *logic) FindFoo() ([]*model.Foo, error) {
	l, err := x.Data.FindFoo()
	if err != nil {
		return l, err
	}

	return l, nil
}

func (x *logic) GetFoo(id string) (*model.Foo, error) {
	r, err := x.Data.GetFoo(id)
	if err != nil {
		return r, err
	}

	return r, nil
}

func (x *logic) CreateFoo(f *model.Foo) (*model.Foo, error) {
	// validation
	if f.Title == "" {
		return f, errors.New("title is required")
	}

	// modification
	f.Title = strings.ToUpper(f.Title)

	// write
	r, err := x.Data.CreateFoo(f)
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
