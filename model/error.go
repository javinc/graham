package model

import "fmt"

// Error model
type Error struct {
	Code    string      `json:"code,omitempty"`
	Message string      `json:"message,omitempty"`
	Panic   bool        `json:"panic,omitempty"`
	Detail  interface{} `json:"detail,omitempty"`
}

func (e Error) Error() string {
	return fmt.Sprintf("%s: %s", e.Code, e.Message)
}
