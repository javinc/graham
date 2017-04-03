package model

import "time"

// Foo test model
type Foo struct {
	ID          string     `json:"id,omitempty" field:"string"`
	UserID      string     `json:"user_id,omitempty" field:"string"`
	Title       string     `json:"title,omitempty" field:"string,required"`
	Description string     `json:"description,omitempty" field:"string"`
	Age         int        `json:"age,omitempty" field:"number,indexed"`
	Taken       *bool      `json:"taken,omitempty" field:"boolean"`
	CreatedAt   *time.Time `json:"created_at,omitempty"`
	UpdatedAt   *time.Time `json:"updated_at,omitempty"`
}

// FooOpts test option model
type FooOpts struct {
	Order  string
	Slice  string
	Field  string
	Filter map[string]interface{}
}
