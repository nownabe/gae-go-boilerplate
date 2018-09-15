package generichandler

import (
	"fmt"
	"net/http"
)

type PingPongHandler struct{}

func NewPingPongHandler() *PingPongHandler {
	return &PingPongHandler{}
}

func (h *PingPongHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "pong")
}
