/**
 * @Author QG
 * @Date  2025/1/16 22:18
 * @description
**/

package gee_web

import (
	"fmt"
	"net/http"
)

// HandlerFunc defines the request handler used by gee
type HandlerFunc func(http.ResponseWriter, *http.Request)

// Engine implement the interface of ServeHTTP
type Engine struct {
	router map[string]HandlerFunc
}

// New is the constructor an instance of gee.Engine
func New() *Engine {
	return &Engine{router: make(map[string]HandlerFunc)}
}

// addRoute defines the method to add route
func (e *Engine) addRoute(method string, pattern string, handler HandlerFunc) {
	key := method + "-" + pattern
	e.router[key] = handler
}

// GET defines the method to add GET route
func (e *Engine) GET(pattern string, handler HandlerFunc) {
	e.addRoute("GET", pattern, handler)
}

// POST defines the method to add POST route
func (e *Engine) POST(pattern string, handler HandlerFunc) {
	e.addRoute("POST", pattern, handler)
}

func (e *Engine) Run(addr string, handler HandlerFunc) (err error) {
	return http.ListenAndServe(addr, e)
}

// ServeHTTP implements the interface of http.Handler
func (e *Engine) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	key := request.Method + "-" + request.URL.Path

	if handler, ok := e.router[key]; !ok {
		fmt.Fprintf(writer, "404 NOT FOUND: %s\n", request.URL)
	} else {
		handler(writer, request)
	}
}
