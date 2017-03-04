package store

import (
	"github.com/gorethink/gorethink"
	db "github.com/javinc/mango/database/rethink"

	"github.com/javinc/graham/model"
)

func init() {
	db.Connect(db.Config{
		Host:    "localhost:28015",
		Db:      "graham",
		MaxOpen: 100,
	})

	db.CreateTable("foo")
}

func (x *store) FindFoo() ([]*model.Foo, error) {
	res, _ := gorethink.Table("foo").Run(db.GetSession())

	l := []*model.Foo{}

	err := res.All(&l)
	if err != nil {
		return l, err
	}

	return l, nil
}

func (x *store) FindFoo1() ([]*model.Foo, error) {
	l := []*model.Foo{}

	r := new(model.Foo)
	r.Title = "list Title"
	r.Description = "list Description"

	l = append(l, r)
	l = append(l, r)
	l = append(l, r)

	return l, nil
}

func (x *store) GetFoo(id string) (*model.Foo, error) {
	r := new(model.Foo)
	r.Title = "detail title"
	r.Description = "detail Description"

	return r, nil
}

func (x *store) CreateFoo(r *model.Foo) (*model.Foo, error) {
	// assuming its inserted to database
	r.ID = "200"

	return r, nil
}
