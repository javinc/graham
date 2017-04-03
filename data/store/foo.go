package store

import (
	"time"

	db "github.com/gorethink/gorethink"
	"github.com/imdario/mergo"

	"github.com/javinc/graham/data/rethink"
	"github.com/javinc/graham/data/util"
	"github.com/javinc/graham/model"
)

const (
	fooTableName = "foo"

	fooErrFind        = "DATA_FOO_FIND"
	fooErrFindOne     = "DATA_FOO_FIND_ONE"
	fooErrGet         = "DATA_FOO_GET"
	fooErrCreate      = "DATA_FOO_CREATE"
	fooErrUpdate      = "DATA_FOO_UPDATE"
	fooErrUpdateCheck = "DATA_FOO_UPDATE_CHECK"
	fooErrRemove      = "DATA_FOO_REMOVE"
	fooErrRemoveCheck = "DATA_FOO_REMOVE_CHECK"
)

func init() {
	rethink.CreateTable(fooTableName)
}

func (x *store) FindFoo(o *model.FooOpts) ([]*model.Foo, error) {
	r := []*model.Foo{}
	// build query
	q := db.Table(fooTableName)
	// build query options
	q = buildOpts(q, o)

	err := rethink.Find(q, &r)
	if err != nil {
		return r, &model.Error{
			Name:    fooErrFind,
			Message: err.Error(),
		}
	}

	return r, nil
}

func (x *store) FindOneFoo() ([]*model.Foo, error) {
	r := []*model.Foo{}
	err := rethink.FindOne(fooTableName, &r)
	if err != nil {
		return r, &model.Error{
			Name:    fooErrFindOne,
			Message: err.Error(),
		}
	}

	return r, nil
}

func (x *store) GetFoo(id string) (*model.Foo, error) {
	r := new(model.Foo)
	err := rethink.Get(fooTableName, id, &r)
	if err != nil {
		return r, &model.Error{
			Name:    fooErrGet,
			Message: err.Error(),
		}
	}

	return r, nil
}

func (x *store) CreateFoo(p *model.Foo) (*model.Foo, error) {
	// meta
	t := time.Now()
	p.CreatedAt = &t
	p.UpdatedAt = &t

	id, err := rethink.Create(fooTableName, p)
	if err != nil {
		return p, &model.Error{
			Name:    fooErrCreate,
			Message: err.Error(),
		}
	}

	p.ID = id

	return p, nil
}

func (x *store) UpdateFoo(p *model.Foo) (*model.Foo, error) {
	r, _ := x.GetFoo(p.ID)
	if r.ID == "" {
		return r, &model.Error{
			Name:    fooErrUpdateCheck,
			Message: "record does not exist",
		}
	}

	// meta update
	t := time.Now()
	p.UpdatedAt = &t

	id := p.ID
	p.ID = ""
	err := rethink.Update(fooTableName, id, p)
	if err != nil {
		return r, &model.Error{
			Name:    fooErrUpdate,
			Message: err.Error(),
		}
	}

	// merge old values with the new
	mergo.MergeWithOverwrite(r, p)

	return r, nil
}

func (x *store) RemoveFoo(id string) (*model.Foo, error) {
	r, _ := x.GetFoo(id)
	if r.ID == "" {
		return r, &model.Error{
			Name:    fooErrRemoveCheck,
			Message: "record does not exist",
		}
	}

	err := rethink.Remove(fooTableName, id)
	if err != nil {
		return r, &model.Error{
			Name:    fooErrRemove,
			Message: err.Error(),
		}
	}

	return r, nil
}

func buildOpts(q db.Term, o *model.FooOpts) db.Term {
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
