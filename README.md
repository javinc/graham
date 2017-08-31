# graham
[![Build Status](https://circleci.com/gh/javinc/mango.svg?style=shield&circle-token=607278cc890cea8c92e97be98eee9b1748c7f75c)](https://circleci.com/gh/javinc/graham)

made with ManGo on top

## usage
### setup
- run server
    - start up rethinkdb server
    - `cd cmd`
    - `go run api.go`

### consume
- listing
    - `curl -X GET localhost:8000/foo`
- detail
    - `curl -X GET localhost:8000/foo/1`
- create
    - `curl -X POST -d '{"title":"dragon","description":"some dragon"}' localhost:8000/foo`
- update
    - `curl -X PATCH -d '{"title":"dragon","description":"some dragon"}' localhost:8000/foo/1`
- remove
    - `curl -X DELETE localhost:8000/foo/1`

### options
- filters
    - `/foo?filter.title=dragon` return row with an id is 1
- sorting
    - `/foo?order=title` ascending by default
    - `/foo?order=title,asc` ascending order
    - `/foo?order=title,desc` descending order
- pagination
    - `/foo?slice=0,10` first 10 rows
- fields
    - `/foo?fields=id`
    - `/foo?fields=id,title`
