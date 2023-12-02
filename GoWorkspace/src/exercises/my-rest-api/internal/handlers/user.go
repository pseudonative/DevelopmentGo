package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/pseudonative/my-rest-api/internal/services"
	"github.com/pseudonative/my-rest-api/pkg/models"
)

type UserHandler struct {
	UserService *services.UserService
}

func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	id, err := h.UserService.CreateUser(r.Context(), user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]int{"id": id})
}

func (h *UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	user, err := h.UserService.GetUser(r.Context(), id)
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(user)
}

func (h *UserHandler) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	users, err := h.UserService.GetAllUsers(r.Context())
	if err != nil {
		http.Error(w, "Failed to retriev users", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(users)
}

func (h *UserHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Failed to update user", http.StatusInternalServerError)
		return
	}

	if err := h.UserService.UpdateUser(r.Context(), user); err != nil {
		http.Error(w, "Failed to update user", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (h *UserHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		http.Error(w, "Invalid User ID", http.StatusBadRequest)
	}

	if err := h.UserService.DeleteUser(r.Context(), id); err != nil {
		http.Error(w, "Failed to delete user", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
