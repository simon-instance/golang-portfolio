package handlers

import (
	"net/http"
	"regexp"

	"github.com/go-chi/chi"
)

type route struct {
	pattern *regexp.Regexp
	handler http.Handler
}

// RegexpHandler handles regex routes
type RegexpHandler struct {
	routes    []*route
	SubRouter *chi.Mux
}

// NewRegexpRouter returns a new regexp handler
func NewRegexpRouter() *RegexpHandler {
	return &RegexpHandler{}
}

// Handler handles the url
func (h *RegexpHandler) Handler(pattern *regexp.Regexp, handler http.Handler) {
	h.routes = append(h.routes, &route{pattern, handler})
}

// HandleFunc handles the url and instantly executes it
func (h *RegexpHandler) HandleFunc(pattern *regexp.Regexp, handler func(http.ResponseWriter, *http.Request)) {
	route := &route{pattern, http.HandlerFunc(handler)}
	h.routes = append(h.routes, route)
}

// ServeHTTP serves to the client
func (h *RegexpHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	for _, route := range h.routes {
		if route.pattern.MatchString(r.URL.Path) {
			route.handler.ServeHTTP(w, r)
			return
		}
	}

	// no pattern matched; send 404 response
	http.NotFound(w, r)
}
