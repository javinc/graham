package store

import (
	rethink "github.com/gorethink/gorethink"
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

// Find base find query
func Find(table string, result interface{}) error {
	return db.Find(rethink.Table(tableName), &result)
}

// FindOne base findone query
func FindOne(table string, result interface{}) error {
	return db.FindOne(rethink.Table(tableName), &result)
}

// Get base get query
func Get(table, id string, result interface{}) error {
	return db.FindOne(rethink.Table(tableName).Get(id), &result)
}
