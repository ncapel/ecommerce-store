package routes

import (
	"fmt"
	"net/http"
)

func routes() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handleNewUser)

}

func handleNewUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World")
}
