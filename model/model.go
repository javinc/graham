/*
  DESIGN RULE
  Package model provides schema of individual application's entities
*/

package model

var (
	trueVal  = true
	falsePtr = !trueVal

	// TruePtr boolean pointer
	TruePtr = &trueVal
	// FalsePtr boolean pointer
	FalsePtr = &falsePtr
)
