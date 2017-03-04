package data

import (
	"context"
)

const key = "data"

// Setter defines a context that enables setting values.
type Setter interface {
	Set(string, interface{})
}

// FromContext returns the Data associated with this context.
func FromContext(c context.Context) Data {
	return c.Value(key).(Data)
}

// ToContext adds the Data to this context if it supports
// the Setter interface.
func ToContext(c Setter, d Data) {
	c.Set(key, d)
}
