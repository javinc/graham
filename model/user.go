package model

import "time"

// User model
type User struct {
	ID        string     `json:"id,omitempty" field:"string"`
	Type      string     `json:"type,omitempty" field:"string"`
	Email     string     `json:"email,omitempty" field:"string"`
	Active    *bool      `json:"taken,omitempty" field:"boolean"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}
