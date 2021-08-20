# highlights-go

## Why Go?

- Simple without too many concepts
- Compiles fast: easier to go to from an interpreted language, good for short dev cycles
- Built for big systems
- Built to scale
- Built by old folk annoyed by C++ for old folk annoyed by hipster languages ;)

## First Steps

- Getting started: https://golang.org/doc/tutorial/getting-started
- Tour of Go: https://tour.golang.org/welcome/1

## Language

- builtin container types:
  - slices
  - maps
- by value vs. by reference
- interfaces
- composite types
- concurrency

## Standard Library

- documentation: https://pkg.go.dev/std
- web server
- JSON
- sorting

## Database Access

- standard: https://golang.org/doc/tutorial/database-access
- https://github.com/jmoiron/sqlx

## Tooling

- go test
- go test -bench=.
- go fmt ./...
- go vet ./...

## Ideas

- Use https://github.com/gin-gonic/gin to implement a REST API (see also https://golang.org/doc/tutorial/web-service-gin)
- Write a price crawler using sync.WaitGroup and https://github.com/PuerkitoBio/goquery
- Write a caching http file server using sync.Mutex and http.ServeContent (to set the ETag), maybe even gzip-encoding the cache entry already
- Implement a simple shooter using https://github.com/hajimehoshi/ebiten