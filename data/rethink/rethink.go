package rethink

import (
	r "github.com/gorethink/gorethink"
	"github.com/javinc/mango/config"
	"github.com/javinc/mango/database/rethink"
)

// init db connection
func init() {
	rethink.Connect(rethink.Config{
		Host:    config.GetString("rethink.host"),
		Db:      config.GetString("rethink.db"),
		MaxOpen: config.GetInt("rethink.max_open"),
	})
}

// CreateTable create table if not exists
func CreateTable(name string) error {
	return rethink.CreateTable(name)
}

// Find base find query
func Find(q r.Term, result interface{}) error {
	return rethink.Find(q, result)
}

// FindOne base findone query
func FindOne(table string, result interface{}) error {
	return rethink.FindOne(r.Table(table), result)
}

// Get base get query
func Get(table, id string, result interface{}) error {
	return rethink.FindOne(r.Table(table).Get(id), result)
}

// Create base create query
func Create(table string, input interface{}) (string, error) {
	return rethink.Create(r.Table(table).Insert(input))
}

// Update base create query
func Update(table, id string, input interface{}) error {
	// modifying without checking if the record exists is fine with rethinkDB
	// update mechanism it will skip the non existent record
	return rethink.Update(r.Table(table).Get(id).Update(input))
}

// Remove base create query
func Remove(table, id string) error {
	return rethink.Remove(r.Table(table).Get(id).Delete())
}
