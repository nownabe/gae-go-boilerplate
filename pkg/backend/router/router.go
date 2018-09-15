package router

import (
	"net/http"

	"github.com/nownabe/gae-go-boilerplate/pkg/generichandler"
)

// Router maps requests to handlers.
type Router struct {
	http.Handler
}

// New returns new Router.
func New() (*Router, error) {
	mux := http.NewServeMux()
	mux.Handle("/ping", generichandler.NewPingPongHandler())
	mux.Handle("/", http.NotFoundHandler())
	return &Router{mux}, nil
}
