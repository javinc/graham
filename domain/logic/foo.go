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

	// data write
	r, err := x.Data.CreateFoo(f)
	if err != nil {
		return r, err
	}

	return r, nil
}
