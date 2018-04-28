package router

import (
	"fmt"
	"net/http"
)

type Router struct {
	routers map[string]http.Handler
}

func (router *Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Called")
}

func New() *Router {
	return &Router{}
}
