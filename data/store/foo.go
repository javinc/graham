package store

import (
	"time"

	db "github.com/gorethink/gorethink"
	"github.com/imdario/mergo"
	"github.com/javinc/mango/errors"

	"github.com/javinc/graham/data/rethink"
	"github.com/javinc/graham/data/util"
	"github.com/javinc/graham/model"
)

const (
	fooTableName = "foo"
)

func init() {
	rethink.CreateTable(fooTableName)
}

func (x *store) FindFoo(o *model.FooOpts) ([]*model.Foo, error) {
	r := []*model.Foo{}

	// build query
	q := db.Table(fooTableName)

	// set owner
	o.Filter["user_id"] = x.User.ID

	// build query options
	q = buildFooOpts(q, o)

	err := rethink.Find(q, &r)
	if err != nil {
		return r, errors.NewError("STORE_FOO_FIND", err)
	}

	return r, nil
}

func (x *store) GetFoo(id string) (*model.Foo, error) {
	r := new(model.Foo)
	err := rethink.Get(fooTableName, id, &r)
	if err != nil {
		return r, errors.NewError("STORE_FOO_GET", err)
	}

	// check owner
	if r.UserID != x.User.ID {
		return new(model.Foo), errors.
			New("STORE_FOO_GET_OWNER", "user record not found")
	}

	return r, nil
}

func (x *store) CreateFoo(p *model.Foo) (*model.Foo, error) {
	// meta
	t := time.Now()
	p.CreatedAt = &t
	p.UpdatedAt = &t

	// set owner
	p.UserID = x.User.ID

	id, err := rethink.Create(fooTableName, p)
	if err != nil {
		return p, errors.NewError("STORE_FOO_CREATE", err)
	}

	p.ID = id

	return p, nil
}

func (x *store) UpdateFoo(p *model.Foo) (*model.Foo, error) {
	r, _ := x.GetFoo(p.ID)
	if r.ID == "" {
		return r, errors.
			New("STORE_FOO_UPDATE_CHK", "record does not exist")
	}

	// meta update
	t := time.Now()
	p.UpdatedAt = &t

	// merge old values with the new
	mergo.MergeWithOverwrite(r, p)

	err := rethink.Update(fooTableName, r.ID, r)
	if err != nil {
		return r, errors.NewError("STORE_FOO_UPDATE", err)
	}

	return r, nil
}

func (x *store) RemoveFoo(id string) (*model.Foo, error) {
	r, _ := x.GetFoo(id)
	if r.ID == "" {
		return r, errors.
			New("STORE_FOO_REMOVE_CHK", "record does not exist")
	}

	err := rethink.Remove(fooTableName, id)
	if err != nil {
		return r, errors.NewError("STORE_FOO_REMOVE", err)
	}

	return r, nil
}

func buildFooOpts(q db.Term, o *model.FooOpts) db.Term {
	// filter
	if len(o.Filter) != 0 {
		q = q.Filter(o.Filter)
	}

	// sort
	if o.Order != "" {
		q = q.OrderBy(util.ParseOptOrder(o.Order))
	}

	// slice
	if o.Slice != "" {
		q = q.Slice(util.ParseOptSlice(o.Slice))
	}

	// pluck
	if o.Field != "" {
		q = q.Pluck(util.ParseOptField(o.Field))
	}

	return q
}
