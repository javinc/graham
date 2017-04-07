package store

import (
	"time"

	db "github.com/gorethink/gorethink"
	"github.com/imdario/mergo"
	"github.com/javinc/mango/database/rethink/util"
	"github.com/javinc/mango/errors"

	"github.com/javinc/graham/data/rethink"
	"github.com/javinc/graham/model"
)

const (
	userTableName = "user"
)

func init() {
	rethink.CreateTable(userTableName)
}

func (x *store) FindUser(o *model.UserOpts) ([]*model.User, error) {
	r := []*model.User{}

	// build query
	q := db.Table(userTableName)

	// build query options
	q = buildUserOpts(q, o)

	err := rethink.Find(q, &r)
	if err != nil {
		return r, errors.NewError("STORE_USER_FIND", err)
	}

	return r, nil
}

func (x *store) GetUser(id string) (*model.User, error) {
	r := new(model.User)
	err := rethink.Get(userTableName, id, &r)
	if err != nil {
		return r, errors.NewError("STORE_USER_GET", err)
	}

	return r, nil
}

func (x *store) CreateUser(p *model.User) (*model.User, error) {
	// meta
	t := time.Now()
	p.CreatedAt = &t
	p.UpdatedAt = &t

	id, err := rethink.Create(userTableName, p)
	if err != nil {
		return p, errors.NewError("STORE_USER_CREATE", err)
	}

	p.ID = id

	return p, nil
}

func (x *store) UpdateUser(p *model.User) (*model.User, error) {
	r, _ := x.GetUser(p.ID)
	if r.ID == "" {
		return r, errors.
			New("STORE_USER_UPDATE_CHK", "record does not exist")
	}

	// meta update
	t := time.Now()
	p.UpdatedAt = &t

	// merge old values with the new
	mergo.MergeWithOverwrite(r, p)

	err := rethink.Update(userTableName, r.ID, r)
	if err != nil {
		return r, errors.NewError("STORE_USER_UPDATE", err)
	}

	return r, nil
}

func (x *store) RemoveUser(id string) (*model.User, error) {
	r, _ := x.GetUser(id)
	if r.ID == "" {
		return r, errors.
			New("STORE_USER_REMOVE_CHK", "record does not exist")
	}

	err := rethink.Remove(userTableName, id)
	if err != nil {
		return r, errors.NewError("STORE_USER_REMOVE", err)
	}

	return r, nil
}

func buildUserOpts(q db.Term, o *model.UserOpts) db.Term {
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
