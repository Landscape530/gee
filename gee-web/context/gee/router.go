package gee

import (
	"log"
	"net/http"
)

type router struct {
	handlers map[string]HandlerFunc
}

func newRouter() *router {
	return &router{handlers: make(map[string]HandlerFunc)}
}

func (rr *router) addRoute(method string, pattern string, handler HandlerFunc) {
	log.Printf("Route router %4s - %s", method, pattern)
	key := method + "-" + pattern
	rr.handlers[key] = handler
}

func (rr *router) handle(c *Context) {
	key := c.Method + "-" + c.Path
	if handler, ok := rr.handlers[key]; ok {
		handler(c)
	} else {
		c.String(http.StatusNotFound, "404 NOT FOUND: %s\n day2-router-handle", c.Path)
	}
}
