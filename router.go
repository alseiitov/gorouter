package gorouter

import (
	"fmt"
	"net/http"
	"regexp"
	"strings"
)

type Handler func(*Context)

type Route struct {
	Pattern *regexp.Regexp
	Handler Handler
	Method  string
}

type Router struct {
	Routes       []Route
	DefaultRoute Handler
}

func NewRouter() *Router {
	return &Router{}
}

func (r *Router) GET(pattern string, handler Handler) {
	r.handle(pattern, handler, http.MethodGet)
}

func (r *Router) POST(pattern string, handler Handler) {
	r.handle(pattern, handler, http.MethodPost)
}

func (r *Router) DELETE(pattern string, handler Handler) {
	r.handle(pattern, handler, http.MethodDelete)
}

func (r *Router) PATCH(pattern string, handler Handler) {
	r.handle(pattern, handler, http.MethodPatch)
}

func (r *Router) handle(pattern string, handler Handler, method string) {
	re := patternToRegex(pattern)
	route := Route{Pattern: re, Handler: handler, Method: method}

	r.Routes = append(r.Routes, route)
}

func patternToRegex(pattern string) *regexp.Regexp {
	splited := strings.Split(pattern, "/")
	for i, v := range splited {
		if strings.HasPrefix(v, ":") {
			splited[i] = `([\w\._-]+)`
		}
	}
	regexStr := fmt.Sprintf("^%s$", strings.Join(splited, "/"))
	return regexp.MustCompile(regexStr)
}

func (r *Router) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	ctx := &Context{Request: request, ResponseWriter: writer}

	for _, rt := range r.Routes {

		if matches := rt.Pattern.FindStringSubmatch(ctx.URL.Path); len(matches) > 0 {
			if request.Method != rt.Method {
				ctx.WriteError(http.StatusMethodNotAllowed, "Method not allowed")
				return
			}

			if len(matches) > 1 {
				ctx.Params = matches[1:]
			}

			rt.Handler(ctx)
			return
		}
	}

	ctx.WriteError(http.StatusNotFound, "404 Not Found")
}
