package model

import "fmt"

// Error model
type Error struct {
	Name    string      `json:"name,omitempty"`
	Message string      `json:"message,omitempty"`
	Panic   bool        `json:"panic,omitempty"`
	Detail  interface{} `json:"detail,omitempty"`
}

func (e Error) Error() string {
	return fmt.Sprintf("%s: %s", e.Name, e.Message)
}
