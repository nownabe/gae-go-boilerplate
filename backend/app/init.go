package main

import (
	"net/http"

	"github.com/nownabe/gae-go-boilerplate/pkg/backend/router"
)

func init() {
	h, err := router.New()
	if err != nil {
		panic(err)
	}
	http.Handle("/", h)
}
