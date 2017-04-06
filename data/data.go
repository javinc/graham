package data

import (
	"context"

	"github.com/javinc/graham/model"
)

// Data package definition
type Data interface {
	FindFoo(o *model.FooOpts) ([]*model.Foo, error)
	GetFoo(id string) (*model.Foo, error)
	CreateFoo(p *model.Foo) (*model.Foo, error)
	UpdateFoo(p *model.Foo) (*model.Foo, error)
	RemoveFoo(id string) (*model.Foo, error)

	FindUser(o *model.UserOpts) ([]*model.User, error)
	GetUser(id string) (*model.User, error)
	CreateUser(p *model.User) (*model.User, error)
	UpdateUser(p *model.User) (*model.User, error)
	RemoveUser(id string) (*model.User, error)
}

// FindFoo returns a list of Foo
func FindFoo(c context.Context, o *model.FooOpts) ([]*model.Foo, error) {
	return FromContext(c).FindFoo(o)
}

// GetFoo returns a detail of Foo
func GetFoo(c context.Context, id string) (*model.Foo, error) {
	return FromContext(c).GetFoo(id)
}

// CreateFoo create a Foo
func CreateFoo(c context.Context, p *model.Foo) (*model.Foo, error) {
	return FromContext(c).CreateFoo(p)
}

// UpdateFoo update a Foo
func UpdateFoo(c context.Context, p *model.Foo) (*model.Foo, error) {
	return FromContext(c).UpdateFoo(p)
}

// RemoveFoo remove a Foo
func RemoveFoo(c context.Context, id string) (*model.Foo, error) {
	return FromContext(c).RemoveFoo(id)
}

// FindUser returns a list of User
func FindUser(c context.Context, o *model.UserOpts) ([]*model.User, error) {
	return FromContext(c).FindUser(o)
}

// GetUser returns a detail of User
func GetUser(c context.Context, id string) (*model.User, error) {
	return FromContext(c).GetUser(id)
}

// CreateUser create a User
func CreateUser(c context.Context, p *model.User) (*model.User, error) {
	return FromContext(c).CreateUser(p)
}

// UpdateUser update a User
func UpdateUser(c context.Context, p *model.User) (*model.User, error) {
	return FromContext(c).UpdateUser(p)
}

// RemoveUser remove a User
func RemoveUser(c context.Context, id string) (*model.User, error) {
	return FromContext(c).RemoveUser(id)
}
