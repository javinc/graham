/*
Package endpoint handles requests and responses over the web and checks for
authenticated requests

DESIGN RULE: Package endpoint provides service of /domain from outside world

Endpoint SHOULD ONLY knows internal packages
[1] /domain

Endpoint SHOULD ONLY do
[1] parse input data and pass to /domain
[2] output formatted results from /domain
*/

package endpoint
