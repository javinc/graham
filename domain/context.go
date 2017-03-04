package domain

import (
	"context"
)

const key = "domain"

// Setter defines a context that enables setting values.
type Setter interface {
	Set(string, interface{})
}

// FromContext returns the Domain associated with this context.
func FromContext(c context.Context) Domain {
	return c.Value(key).(Domain)
}

// ToContext adds the Domain to this context if it supports
// the Setter interface.
func ToContext(c Setter, d Domain) {
	c.Set(key, d)
}
