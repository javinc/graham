package model

import "time"

// User model
type User struct {
	ID        string     `data:"id,omitempty" json:"id,omitempty"     field:"string"`
	Type      string     `data:"type"         json:"type,omitempty"   field:"string"`
	Name      string     `data:"name"         json:"name,omitempty"   field:"string"`
	Email     string     `data:"email"        json:"email,omitempty"  field:"string"`
	Active    bool       `data:"active"       json:"active,omitempty" field:"boolean"`
	CreatedAt *time.Time `data:"created_at"   json:"created_at,omitempty"`
	UpdatedAt *time.Time `data:"updated_at"   json:"updated_at,omitempty"`
}
