package routes

import (
	"net/http"

	"github.com/vidalme/ativia-api/internal/controllers"
)

func RegisterRoutes() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/", controllers.Index)
	mux.HandleFunc("POST /user", controllers.AddUser)
	mux.HandleFunc("POST /user/createWithArray", controllers.AddUsers)
	mux.HandleFunc("GET /user/{userName}", controllers.GetUser)
	mux.HandleFunc("PUT /user/{userName}", controllers.UpdateUser)
	mux.HandleFunc("DELETE /user/{userName}", controllers.DeleteUser)

	return mux
}
