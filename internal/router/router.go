package router

import (
	"github.com/maindotmarcell/http-from-scratch/internal/http"
)

type HandlerFunc func(http.Request) string

type Router struct {
	routes map[string]map[string]HandlerFunc
}

func New() *Router {
	r := &Router{
		routes: make(map[string]map[string]HandlerFunc),
	}
	r.routes["GET"] = make(map[string]HandlerFunc)
	r.routes["POST"] = make(map[string]HandlerFunc)

	return r
}

func (r *Router) HandleGet(path string, handlerFunc HandlerFunc) {
	r.routes["GET"][path] = handlerFunc
}

func (r *Router) HandlePost(path string, handlerFunc HandlerFunc) {
	r.routes["POST"][path] = handlerFunc
}

func (r *Router) Route(req http.Request) HandlerFunc {
	return r.routes[req.RequestLine.Method][req.RequestLine.Path]
}
