package routes

import (
	"net/http"
)

func InitRoutes() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("POST /users", handleNewUser)
	return mux
}
