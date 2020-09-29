package router

import "net/http"

// Router is an interface implementing methods for the server
type Router interface {
	GET(uri string, f func(res http.ResponseWriter, req *http.Request))
	POST(uri string, f func(res http.ResponseWriter, req *http.Request))
	SERVE(port string)
}
