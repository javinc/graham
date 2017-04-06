package model

import "time"

// User model
type User struct {
	ID        string     `data:"id,omitempty" json:"id,omitempty"          field:"string"`
	Name      string     `data:"name"         json:"name,omitempty"        field:"string"`
	Email     string     `data:"email"        json:"email,omitempty"       field:"string"`
	Password  string     `data:"password"     json:"password,omitempty"    field:"string"`
	Verified  bool       `data:"verified"     json:"verified,omitempty"    field:"boolean"`
	CreatedAt *time.Time `data:"created_at"   json:"created_at,omitempty"`
	UpdatedAt *time.Time `data:"updated_at"   json:"updated_at,omitempty"`
}

// UserOpts option model
type UserOpts struct {
	Order  string
	Slice  string
	Field  string
	Filter map[string]interface{}
}
