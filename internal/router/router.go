package router

import (
	"github.com/maindotmarcell/http-from-scratch/internal/http"
)

type HandlerFunc func(http.Request) string

type Router struct {
	routes map[string]map[string]HandlerFunc
}

// Initializes a new Router struct with GET and POST keys populated
func New() *Router {
	r := &Router{
		routes: make(map[string]map[string]HandlerFunc),
	}
	r.routes["GET"] = make(map[string]HandlerFunc)
	r.routes["POST"] = make(map[string]HandlerFunc)

	return r
}

// Handles GET requests
func (r *Router) HandleGet(path string, handlerFunc HandlerFunc) {
	r.routes["GET"][path] = handlerFunc
}

// Handles POST requests
func (r *Router) HandlePost(path string, handlerFunc HandlerFunc) {
	r.routes["POST"][path] = handlerFunc
}

// Finds the handler function for the method and path present in the request
// Returns nil if not found
func (r *Router) Route(req http.Request) HandlerFunc {
	return r.routes[req.RequestLine.Method][req.RequestLine.Path]
}
