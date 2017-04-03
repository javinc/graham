/*
Package domain process data and validates them before going to database, also
implments business logic

DESIGN RULE: Package domain process and check data base on business rules

Domain SHOULD ONLY knows internal packages
    [1] /data
    [1] /platform

Domain SHOULD ONLY do
    [1] business or domain logic and use cases
    [2] access records from /data
    [3] use third party service from /platform
*/

package domain
