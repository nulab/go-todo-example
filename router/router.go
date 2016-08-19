package router

import (
	"net/http"
	"regexp"
)

type Route struct {
	Pattern    *regexp.Regexp
	Handler    http.Handler
	HTTPMethod string
}

type Router struct {
	routes []*Route
}


func (r *Router) HandleFunc(pattern string, httpMethod string, f func(http.ResponseWriter, *http.Request)) {
	r.routes = append(r.routes, &Route{
		HTTPMethod:httpMethod,
		Pattern: regexp.MustCompile(pattern + "$"),
		Handler: http.HandlerFunc(f),
	})
}

func (rer *Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	for _, route := range rer.routes {
		if route.HTTPMethod == r.Method && route.Pattern.MatchString(r.URL.Path) {
			route.Handler.ServeHTTP(w, r)
			return
		}
	}
	http.NotFound(w, r)
}