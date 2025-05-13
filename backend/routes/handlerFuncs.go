package routes

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ncapel/ecommerce-store/config"
	"github.com/ncapel/ecommerce-store/models"
)

func handleNewUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	if user.Name == "" || user.Email == "" || user.Password == "" {
		http.Error(w, "Missing required fields", http.StatusBadRequest)
		return
	}

	id, err := models.CreateUser(config.Db, user.Name, user.Password, user.Email)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error creating user: %v", err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "User created with ID: %d", id)
}

func handleDelUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	if user.ID < 0 || user.ID > 2147483647 {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	id, _, err := models.DeleteUser(config.Db, user)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error deleting user: %v", err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "User deleted with ID: %d", id)
}
