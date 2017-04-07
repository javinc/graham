package logic

import (
	"github.com/javinc/mango/errors"
	"github.com/javinc/mango/server/auth"

	"github.com/javinc/graham/model"
)

const (
	userAuthErrKey = "DOMAIN_USER_AUTH_"
)

func (x *logic) RegisterUser(p *model.User) (*model.User, error) {
	u, err := x.CreateUser(p)
	if err != nil {
		return u, err
	}

	// send email verification

	return u, err
}

func (x *logic) LoginUser(email, pass string) (map[string]interface{}, error) {
	m := map[string]interface{}{}

	u, err := x.FindUserByEmail(email)
	if err != nil {
		return m, err
	}

	// check password
	if hash(pass) != u.Password {
		return m, errors.
			New("LOGIC_USER_AUTH_LOGIN", "invalid email or password")
	}

	m["id"] = u.ID
	m["email"] = u.Email

	// generate JWT
	t, err := auth.CreateToken(m)
	if err != nil {
		return m, errors.PanicError("LOGIC_USER_AUTH_JWT", err)
	}

	m["token"] = t
	m["name"] = u.Name

	return m, nil
}

func (x *logic) CurrentUser() (*model.User, error) {
	u := x.User
	if u.ID == "" {
		return u, errors.
			New("LOGIC_USER_AUTH_NO_USER", "current credentails not valid")
	}

	// mask password
	u.Password = "******"

	return u, nil
}
