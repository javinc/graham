package model

import "time"

// Foo test model
type Foo struct {
	ID          string     `data:"id,omitempty" json:"id,omitempty"          field:"string"`
	UserID      string     `data:"user_id"      json:"user_id,omitempty"     field:"string"`
	Title       string     `data:"title"        json:"title,omitempty"       field:"string,required"`
	Description string     `data:"description"  json:"description,omitempty" field:"string"`
	Age         int        `data:"age"          json:"age,omitempty"         field:"number"`
	Taken       bool       `data:"taken"        json:"taken,omitempty"       field:"boolean"`
	CreatedAt   *time.Time `data:"created_at"   json:"created_at,omitempty"`
	UpdatedAt   *time.Time `data:"updated_at"   json:"updated_at,omitempty"`
}

// FooOpts test option model
type FooOpts struct {
	Order  string
	Slice  string
	Field  string
	Filter map[string]interface{}
}
