package store

import "github.com/javinc/graham/model"

func (x *store) FindUser() ([]*model.User, error) {
	l := []*model.User{}

	r := new(model.User)
	r.ID = "100"
	r.Email = "list Email"

	l = append(l, r)
	l = append(l, r)
	l = append(l, r)

	return l, nil
}

func (x *store) GetUser(id string) (*model.User, error) {
	r := new(model.User)
	r.ID = "100"
	r.Type = "client"
	r.Email = "detail Email"

	return r, nil
}
