/*
  DESIGN RULE

  endpoint package SHOULD ONLY do
    [1] pass parsed value
    [2] output result from domain functions

  endpoint SHOULD ONLY knows internal package
    [1] domain
*/

// Package endpoint provides service of domain package rom http request
package endpoint
