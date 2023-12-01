package handlers

import (
	"encoding/json"
	"net/http"

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
