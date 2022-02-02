package mux

import (
	"net/http"
)

type Route struct {
	// Request handler for the route.
	handler http.Handler
	// If true, this route never matches: it is only used to build URLs.
	buildOnly bool
	// The name used to build URLs.
	name string
	// Error resulted from building a route.
	err error

	// "global" reference to all named routes
	namedRoutes map[string]*Route

	// config possibly passed in from `Router`
}

type Router struct {
	namedRoutes map[string]*Route
}

func NewRouter() *Router {
	return &Router{namedRoutes: make(map[string]*Route)}
}

func (r *Router) StrictSlash(value bool) *Router {
	r.strictSlash = value
	return r
}

func (r *Router) GetRoute(name string) *Route {
	return r.namedRoutes[name]
}