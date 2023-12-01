package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/pseudonative/my-rest-api/internal/services"
)

type UserHandler struct {
	UserService *services.UserService
}

func (h *UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(r.URL.Query().Get("id"))
	user, err := h.UserService.GetUser(r.Context(), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(user)
}
