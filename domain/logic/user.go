package logic

import "github.com/javinc/graham/model"

func (x *logic) FindUser() ([]*model.User, error) {
	return x.Data.FindUser()
}

func (x *logic) GetUser(id string) (*model.User, error) {
	return x.Data.GetUser(id)
}
