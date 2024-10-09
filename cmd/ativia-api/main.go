package main

import (
	"net/http"

	"github.com/vidalme/ativia-api/internal/routes"
)

func main() {
	mux := routes.RegisterRoutes()
	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.ListenAndServe(":2113", mux)
}
