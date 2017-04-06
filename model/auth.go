package model

import "time"

// Auth app model use for accessing private resource
// mainly use on JWT as payload model
type Auth struct {
	// ID refers to user ID
	ID string `json:"id,omitempty"      field:"string"`
	// access type
	Access     string     `json:"access,omitempty"  field:"string"`
	Token      string     `json:"token,omitempty"   field:"string"`
	Expiration *time.Time `json:"expiration,omitempty"`
}
