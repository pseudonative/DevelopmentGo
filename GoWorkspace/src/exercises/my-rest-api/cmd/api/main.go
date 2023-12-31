package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/lib/pq"
	"github.com/pseudonative/my-rest-api/internal/handlers"
	"github.com/pseudonative/my-rest-api/internal/repository"
	"github.com/pseudonative/my-rest-api/internal/services"
)

func main() {
	const (
		host     = "localhost"
		port     = 5432
		user     = "postgres"
		password = "mysecretpassword"
		dbname   = "postgres"
	)

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal("Could not connect to databvase:", err)
	}
	defer db.Close()

	if err = db.Ping(); err != nil {
		log.Fatal("Error pinging the database: ", err)
	}

	userRepo := &repository.UserRepository{DB: db}
	userService := &services.UserService{Repo: userRepo}
	userHandler := &handlers.UserHandler{UserService: userService}

	mux := http.NewServeMux()

	mux.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			userHandler.CreateUser(w, r)
		case http.MethodGet:
			userHandler.GetAllUsers(w, r)
		case http.MethodPut:
			userHandler.UpdateUser(w, r)
		case http.MethodDelete:
			userHandler.DeleteUser(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	mux.HandleFunc("/user", userHandler.GetUser)

	log.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal(err)
	}
}
