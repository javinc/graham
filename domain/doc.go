/*
Package domain process data and validates them before going to database, also
implments business logic

DESIGN RULE: Package domain process and check data base on business rules

Domain SHOULD ONLY knows internal packages
    - /data
    - /platform

Domain SHOULD ONLY do
    - business or domain logic and use cases
    - access records from /data
    - use third party service from /platform
*/

package domain
