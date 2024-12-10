package api

import (
	"net/http"
)

func RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/ping", PingHandler)
	mux.HandleFunc("/user", CreateUserHandler)
	mux.HandleFunc("/", PingHandler)
}