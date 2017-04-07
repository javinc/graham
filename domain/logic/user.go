package logic

import (
	"crypto/md5"
	"encoding/hex"

	"github.com/javinc/mango/errors"

	"github.com/javinc/graham/model"
)

func (x *logic) FindUser(o *model.UserOpts) ([]*model.User, error) {
	r, err := x.Data.FindUser(o)
	if err != nil {
		return r, err
	}

	return r, nil
}

func (x *logic) FindOneUser(o *model.UserOpts) (*model.User, error) {
	d := new(model.User)
	r, err := x.FindUser(o)
	if err != nil {
		return d, err
	}

	if len(r) == 0 {
		return d, errors.New("LOGIC_USER_FIND1", "record not found")
	}

	return r[0], nil
}

func (x *logic) GetUser(id string) (*model.User, error) {
	// validation
	if id == "" {
		return new(model.User), errors.
			New("LOGIC_USER_GET_CHK", "id param is required")
	}

	r, err := x.Data.GetUser(id)
	if err != nil {
		return r, err
	}

	return r, nil
}

func (x *logic) CreateUser(p *model.User) (*model.User, error) {
	// validation
	if p.Name == "" || p.Email == "" || p.Password == "" {
		return p, errors.
			New("LOGIC_USER_CREATE_CHK",
				"name, email, and password field are required")
	}

	// password check
	if len(p.Password) < 6 {
		return p, errors.
			New("LOGIC_USER_PASS_LEN",
				"password should have atleast 6 characters")
	}

	// no duplicate email
	u, _ := x.FindUserByEmail(p.Email)
	if u.ID != "" {
		return p, errors.
			New("LOGIC_USER_EMAIL_EXISTS",
				"email already in use")
	}

	// modification
	p.Password = hash(p.Password)

	// write
	r, err := x.Data.CreateUser(p)
	if err != nil {
		return r, err
	}

	return r, nil
}

func (x *logic) UpdateUser(p *model.User) (*model.User, error) {
	// validation
	if p.ID == "" {
		return p, errors.New("LOGIC_USER_UPDATE_CHK", "id field is required")
	}

	// write
	p, err := x.Data.UpdateUser(p)
	if err != nil {
		return p, err
	}

	return p, nil
}

func (x *logic) RemoveUser(id string) (*model.User, error) {
	// validation
	if id == "" {
		return new(model.User), errors.
			New("LOGIC_USER_REMOVE_CHK", "id param is required")
	}

	// write
	r, err := x.Data.RemoveUser(id)
	if err != nil {
		return r, err
	}

	return r, nil
}

func (x *logic) FindUserByEmail(email string) (*model.User, error) {
	return x.FindOneUser(&model.UserOpts{
		Filter: map[string]interface{}{
			"email": email,
		},
	})
}

func hash(s string) string {
	salt := "andpepper"
	hasher := md5.New()
	hasher.Write([]byte(s + salt))

	return hex.EncodeToString(hasher.Sum(nil))
}
