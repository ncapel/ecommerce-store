package routes

import (
	"net/http"
)

func InitRoutes() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("POST /users", handleNewUser)
	mux.HandleFunc("DELETE /users", handleDelUser)
	mux.HandleFunc("PATCH /users/", handlePatchUser)
	return mux
}
