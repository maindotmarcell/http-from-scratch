# HTTP from scratch

This is a custom HTTP server, built using only TCP socket programming. It serves requests by reading from and writing to raw TCP
connections while following the HTTP 1.1 specification. The goal of this projet is to gain a deeper understanding of networking
protocols (namely TCP and HTTP).

## Supported paths

- `GET "/"`

## How to add additional paths

1. Create a handler function in [handler.go](internal/handler/handler.go) such that it takes in an http request and
   returns an http 1.1 response compatible string. Type definitions and util functions for requests and responses
   can be found in [request.go](internal/http/request.go) and [response.go](internal/http/response.go) respectively.
2. Add a new path in [main.go](cmd/main.go) `s.Router.HandleGet("/example-path", handler.HandleExamplePath)`.

## Roadmap

- Keeping connections alive and support for `Connection: close` headers
- Support for common headers
- Creating and fetching files