package Handlers

import (
	"NOW/logic/Entities"
	"NOW/logic/Repositories/user"
	"encoding/json"
	"net/http"
)

type UserHandler struct {
	repo Repositories.UserRepository
}

func NewUserHandler(repo Repositories.UserRepository) *UserHandler {
	return &UserHandler{repo: repo}
}

func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case http.MethodPost:
		var user Entities.User

		// Decode the request body into the user struct
		err := json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// Now you can use the user struct
		err = h.repo.Create(&user)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Send a response back to the client
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(user)

	default:
		// Give an error message.
		http.Error(w, "Invalid method", http.StatusMethodNotAllowed)
	}

}
