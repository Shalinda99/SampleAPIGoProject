package router

import (
	"net/http"
	"task2/controller"
)

type Router struct {
	controllers map[string]controller.Controller
}

func NewRouter() *Router {
	return &Router{
		controllers: make(map[string]controller.Controller),
	}
}

func (r *Router) RegisterController(route string, controller controller.Controller) {
	r.controllers[route] = controller
}

func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	path := req.URL.Path
	if controller, ok := r.controllers[path]; ok {
		controller.HandleRequest(w, req)
	} else {
		http.Error(w, "Not Found", http.StatusNotFound)
	}
}
