package router

import (
	"sort"
	"strings"

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

// Handles assigning handler functions to method and path combinations
func (r *Router) handle(method string, path string, handlerFunc HandlerFunc) {
	r.routes[method][path] = handlerFunc
}

// Handles GET requests. Wrapper function for assigning handler functions to GET paths
func (r *Router) HandleGet(path string, handlerFunc HandlerFunc) {
	r.handle("GET", path, handlerFunc)
}

// Handles POST requests. Wrapper function for assigning handler functions to POST paths.
func (r *Router) HandlePost(path string, handlerFunc HandlerFunc) {
	r.handle("POST", path, handlerFunc)
}

// Finds the handler function for the method and path present in the request.
// If exact path match is not found it will look for the longest matching prefix in the path.
// Returns nil if not found.
func (r *Router) Route(req http.Request) HandlerFunc {
	path := req.RequestLine.Path
	method := req.RequestLine.Method

	// Get all registered paths for this method and sort them by length (longest first)
	var paths []string
	for registeredPath := range r.routes[method] {
		paths = append(paths, registeredPath)
	}
	sort.Slice(paths, func(i, j int) bool {
		return len(paths[i]) > len(paths[j])
	})

	// Check each registered route in order (longest to shortest)
	for _, registeredPath := range paths {
		if strings.HasPrefix(path, registeredPath) {
			return r.routes[method][registeredPath]
		}
	}
	return nil
}
