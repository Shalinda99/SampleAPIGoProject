package controller

import "net/http"

type Controller interface {
	HandleRequest(w http.ResponseWriter, r *http.Request)
}

type DefaultController struct{}

func NewDefaultController() *DefaultController {
	return &DefaultController{}
}

func (c *DefaultController) HandleRequest(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Not Found", http.StatusNotFound)
}
