package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/pseudonative/profile_page/internal/api"
	"github.com/pseudonative/profile_page/internal/login"
)

func main() {
	fs := http.FileServer(http.Dir("."))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.Handle("/", api.TokenAuthMiddleWare(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "profile.html")
	})))
	http.Handle("/login", http.HandlerFunc(login.LoginHandler))
	fmt.Println("Server Started at http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
