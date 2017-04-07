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
	userTableName = "user"

	userErrFind        = "DATA_FOO_FIND"
	userErrFindOne     = "DATA_FOO_FIND_ONE"
	userErrGet         = "DATA_FOO_GET"
	userErrCreate      = "DATA_FOO_CREATE"
	userErrUpdate      = "DATA_FOO_UPDATE"
	userErrUpdateCheck = "DATA_FOO_UPDATE_CHECK"
	userErrRemove      = "DATA_FOO_REMOVE"
	userErrRemoveCheck = "DATA_FOO_REMOVE_CHECK"
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
		return r, &model.Error{
			Name:    userErrFind,
			Message: err.Error(),
		}
	}

	return r, nil
}

func (x *store) GetUser(id string) (*model.User, error) {
	r := new(model.User)
	err := rethink.Get(userTableName, id, &r)
	if err != nil {
		return r, &model.Error{
			Name:    userErrGet,
			Message: err.Error(),
		}
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
		return p, &model.Error{
			Name:    userErrCreate,
			Message: err.Error(),
		}
	}

	p.ID = id

	return p, nil
}

func (x *store) UpdateUser(p *model.User) (*model.User, error) {
	r, _ := x.GetUser(p.ID)
	if r.ID == "" {
		return r, &model.Error{
			Name:    userErrUpdateCheck,
			Message: "record does not exist",
		}
	}

	// meta update
	t := time.Now()
	p.UpdatedAt = &t

	// merge old values with the new
	mergo.MergeWithOverwrite(r, p)

	err := rethink.Update(userTableName, r.ID, r)
	if err != nil {
		return r, &model.Error{
			Name:    userErrUpdate,
			Message: err.Error(),
		}
	}

	return r, nil
}

func (x *store) RemoveUser(id string) (*model.User, error) {
	r, _ := x.GetUser(id)
	if r.ID == "" {
		return r, &model.Error{
			Name:    userErrRemoveCheck,
			Message: "record does not exist",
		}
	}

	err := rethink.Remove(userTableName, id)
	if err != nil {
		return r, &model.Error{
			Name:    userErrRemove,
			Message: err.Error(),
		}
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
