package main

import (
	"database/sql"
	"log"
	"net/http"

	_ "github.com/lib/pq"
	"github.com/pseudonative/my-rest-api/internal/handlers"
	"github.com/pseudonative/my-rest-api/internal/repository"
	"github.com/pseudonative/my-rest-api/internal/services"
)

func main() {
	db, err := sql.Open("postgres", "postres")
	if err != nil {
		log.Fatal("Could not connect to databvase:", err)
	}
	userRepo := &repository.UserRepository{DB: db}
	userService := &services.UserService{Repo: userRepo}
	userHandler := &handlers.UserHandler{UserService: userService}

	mux := http.NewServeMux()
	mux.HandleFunc("/user", userHandler.GetUser)

	log.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal(err)
	}
}
