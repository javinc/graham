package store

import (
	db "github.com/javinc/mango/database/rethink"

	"github.com/javinc/graham/data"
	"github.com/javinc/graham/model"
)

type store struct {
	User *model.User
}

func init() {
	db.Connect(db.Config{
		Host:    "localhost:28015",
		Db:      "graham",
		MaxOpen: 100,
	})

	db.CreateTable("foo")
}

// New create new store object
func New(u *model.User) data.Data {
	d := &store{u}

	return d
}
