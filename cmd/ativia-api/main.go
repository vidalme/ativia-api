package main

import (
	"net/http"

	"github.com/vidalme/ativia-api/internal/routes"
)

func main() {
	routes.RegisterRoutes()
	http.ListenAndServe(":2113", nil)
}
