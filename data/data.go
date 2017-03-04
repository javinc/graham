package data

import (
	"context"

	"github.com/javinc/mango/model"
)

// Data package definition
type Data interface {
	FindFoo() ([]*model.Foo, error)
	GetFoo(id string) (*model.Foo, error)
	CreateFoo(r *model.Foo) (*model.Foo, error)
	// RemoveFoo(id string) (*model.Foo, error)
	// UpdateFoo(r *model.Foo, id string) (*model.Foo, error)

	FindUser() ([]*model.User, error)
	GetUser(id string) (*model.User, error)
}

// FindFoo returns a list of Foo
func FindFoo(c context.Context) ([]*model.Foo, error) {
	return FromContext(c).FindFoo()
}

// CreateFoo returns a detail of Foo
func CreateFoo(c context.Context, r *model.Foo) (*model.Foo, error) {
	return FromContext(c).CreateFoo(r)
}

// GetFoo returns a detail of Foo
func GetFoo(c context.Context, id string) (*model.Foo, error) {
	return FromContext(c).GetFoo(id)
}

// FindUser returns a list of User
func FindUser(c context.Context) ([]*model.User, error) {
	return FromContext(c).FindUser()
}

// GetUser returns a detail of User
func GetUser(c context.Context, id string) (*model.User, error) {
	return FromContext(c).GetUser(id)
}
