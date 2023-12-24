package login

import (
	"net/http"

	"github.com/pseudonative/profile_page/internal/api"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	username := r.URL.Query().Get("username")
	tokenString, err := api.GenerateJWT(username)
	if err != nil {
		http.Error(w, "error generating token", http.StatusInternalServerError)
		return
	}
	w.Write([]byte(tokenString))
}
