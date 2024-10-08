package main

import (
	"net/http"

	"github.com/vidalme/ativia-api/internal/routes"
)

func main() {
	mux := routes.RegisterRoutes()
	http.ListenAndServe(":2113", mux)
}
