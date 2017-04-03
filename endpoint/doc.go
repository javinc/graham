/*
Package endpoint handles requests and responses over the web and checks for
authenticated requests

DESIGN RULE: Package endpoint provides service of /domain from outside world

Endpoint SHOULD ONLY knows internal packages
    - /domain

Endpoint SHOULD ONLY do
    - parse input data and pass to /domain
    - output formatted results from /domain
*/

package endpoint
