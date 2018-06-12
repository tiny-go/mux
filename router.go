package mux

import "net/http"

// Node interface.
type Node interface {
	Add(string, http.Handler) error
	Get(string) (http.Handler, map[string]string, bool)
}

// Router represents simple router with configurable route tree.
type Router struct {
	nf    func() Node
	nodes map[string]Node
}

// NewRouter creates a Router.
func NewRouter(nf func() Node) *Router {
	return &Router{nf, make(map[string]Node)}
}

func (mux *Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	node, ok := mux.nodes[r.Method]
	if !ok {
		http.Error(w, "not found", http.StatusNotFound)
		return
	}
	h, _, ok := node.Get(r.URL.Path)
	if !ok {
		http.Error(w, "not found", http.StatusNotFound)
		return
	}
	h.ServeHTTP(w, r)
}

// Handle registers a new request handle with the given path and method.
func (mux *Router) Handle(method, pattern string, handler http.Handler) error {
	node, ok := mux.nodes[method]
	if !ok {
		node, mux.nodes[method] = mux.nf(), node
	}
	return node.Add(pattern, handler)
}
