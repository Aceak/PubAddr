package server

import (
	"PubAddr/internal/logger"
	"net/http"
)

type Router struct {
	mux *http.ServeMux
}

func NewRouter() *Router {
	logger.Debug("Initializing HTTP router")
	return &Router{mux: http.NewServeMux()}
}

func (r *Router) Handler() http.Handler {
	return r.mux
}

func (r *Router) GET(path string, h http.HandlerFunc, mws ...Middleware) {
	r.register(http.MethodGet, path, h, mws...)
}

func (r *Router) POST(path string, h http.HandlerFunc, mws ...Middleware) {
	r.register(http.MethodPost, path, h, mws...)
}

func (r *Router) register(method, path string, h http.HandlerFunc, mws ...Middleware) {
	final := Use(h, mws...)

	r.mux.HandleFunc(path, func(w http.ResponseWriter, req *http.Request) {
		if req.Method != method {
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
			return
		}
		final.ServeHTTP(w, req)
	})
}
