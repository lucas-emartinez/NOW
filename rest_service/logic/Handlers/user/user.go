package user

import (
	"NOW/rest_service/logic/Repositories/user"
	dbEntity "NOW/rest_service/logic/entities/db"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"regexp"
	"strconv"
	"time"
)

type UserHandler struct {
	repo Repositories.UserRepository
}

var (
	regexNumber  = regexp.MustCompile(`[0-9]`)                         // checks for a number
	regexLower   = regexp.MustCompile(`[a-z]`)                         // checks for a lower case letter
	regexUpper   = regexp.MustCompile(`[A-Z]`)                         // checks for an upper case letter
	regexSpecial = regexp.MustCompile(`[\!\*\@\#\$\%\^\&\*\(\)_\=\+]`) // checks for a special character
)

func NewUserHandler(repo Repositories.UserRepository) *UserHandler {
	return &UserHandler{repo: repo}
}

func validatePassword(password string) error {
	// Validate password length
	if len(password) < 8 || len(password) > 20 {
		return errors.New("Password must be at least 8 characters and less than 20 characters")
	}
	if !regexNumber.MatchString(password) {
		return errors.New("Password must contain at least a number")
	}
	if !regexLower.MatchString(password) {
		return errors.New("Password must contain at least a lowercase letter")
	}
	if !regexUpper.MatchString(password) {
		return errors.New("Password must contain at least an uppercase letter")
	}
	if !regexSpecial.MatchString(password) {
		return errors.New("Password must contain at least a special character")
	}

	return nil
}

func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 2*time.Second)
	defer cancel()

	var user dbEntity.User

	// Decode the request body into the user struct
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Validate password length
	err = validatePassword(user.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// Now you can use the user struct
	err = h.repo.Create(ctx, &user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Send a response back to the client
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}

func (h *UserHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {

	ctx, cancel := context.WithTimeout(r.Context(), 2*time.Second)
	defer cancel()

	params := r.Context().Value("params").(map[string]string)
	dniQuery := params["dni"] // assuming you have a "dni" parameter in your route
	dni, err := strconv.Atoi(dniQuery)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var updatedUser dbEntity.User

	err = json.NewDecoder(r.Body).Decode(&updatedUser)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// fetch the existing user from the database
	existingUser, err := h.repo.GetByDNI(ctx, dni)
	fmt.Println(existingUser)
	if err != nil {
		recover()
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	// update the fields that are allowed to be changed
	existingUser.Name = updatedUser.Name
	existingUser.LastName = updatedUser.LastName
	existingUser.IdentityCheck = updatedUser.IdentityCheck

	err = h.repo.Update(ctx, existingUser)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Send a response back to the client
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(existingUser)
}
func (h *UserHandler) GetByDNI(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 10*time.Second)
	defer cancel()

	params := r.Context().Value("params").(map[string]string)

	dniQuery := params["dni"]
	if dniQuery == "" {
		http.Error(w, "Missing DNI", http.StatusBadRequest)
	}
	dni, err := strconv.Atoi(dniQuery)

	if err != nil {
		http.Error(w, "DNI must be an integer", http.StatusBadRequest)
		return
	}

	user, err := h.repo.GetByDNI(ctx, dni)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	if user == nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(user)
}
func (h *UserHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 10*time.Second)
	defer cancel()

	params := r.Context().Value("params").(map[string]string)

	dniQuery := params["dni"]
	if dniQuery == "" {
		http.Error(w, "Missing DNI", http.StatusBadRequest)
	}
	dni, err := strconv.Atoi(dniQuery)
	if err != nil {
		http.Error(w, "DNI must be an integer", http.StatusBadRequest)
		return
	}

	err = h.repo.Delete(ctx, dni)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Usuario eliminado"))
}
