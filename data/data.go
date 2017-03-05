/*
  DESIGN RULE
	Package data provides access from database

	Data SHOULD ONLY knows internal packages
		[1] /platform

  Data SHOULD ONLY do
    [1] database queries
    [2] meta implementation
*/

package data

import (
	"context"

	"github.com/javinc/graham/model"
)

// Data package definition
type Data interface {
	FindFoo() ([]*model.Foo, error)
	GetFoo(id string) (*model.Foo, error)
	CreateFoo(r *model.Foo) (*model.Foo, error)
	UpdateFoo(r *model.Foo) (*model.Foo, error)
	// RemoveFoo(id string) (*model.Foo, error)

	FindUser() ([]*model.User, error)
	GetUser(id string) (*model.User, error)
}

// FindFoo returns a list of Foo
func FindFoo(c context.Context) ([]*model.Foo, error) {
	return FromContext(c).FindFoo()
}

// GetFoo returns a detail of Foo
func GetFoo(c context.Context, id string) (*model.Foo, error) {
	return FromContext(c).GetFoo(id)
}

// CreateFoo create a Foo
func CreateFoo(c context.Context, r *model.Foo) (*model.Foo, error) {
	return FromContext(c).CreateFoo(r)
}

// UpdateFoo update a Foo
func UpdateFoo(c context.Context, r *model.Foo) (*model.Foo, error) {
	return FromContext(c).UpdateFoo(r)
}

// FindUser returns a list of User
func FindUser(c context.Context) ([]*model.User, error) {
	return FromContext(c).FindUser()
}

// GetUser returns a detail of User
func GetUser(c context.Context, id string) (*model.User, error) {
	return FromContext(c).GetUser(id)
}
