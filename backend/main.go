package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ncapel/ecommerce-store/config"
	"github.com/ncapel/ecommerce-store/models"
)

func main() {
	config.ConnectDB()
	//	config.SeedDb()
	mux := http.NewServeMux()
	mux.HandleFunc("/", handleRoot)

	mux.HandleFunc("POST /users", handleNewUser)

	http.ListenAndServe(":8080", mux)

}

func handleRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World")
}

func handleNewUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if user.Name == "" {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	models.CreateUser(config.Db, user.Name, user.Email, user.Password)

	w.WriteHeader(http.StatusNoContent)
}
