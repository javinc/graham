package model

// Error model
type Error struct {
	Name    string      `json:"name,omitempty" field:"string"`
	Message string      `json:"message,omitempty" field:"string"`
	Detail  interface{} `json:"detail,omitempty" field:"string"`
}
