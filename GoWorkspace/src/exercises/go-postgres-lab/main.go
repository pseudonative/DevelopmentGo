package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

func main() {
	db := connectDB()
	defer db.Close()
	createTables(db)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello from go with postgres")
	})
	http.ListenAndServe(":8080", nil)

}

func connectDB() *sql.DB {
	connStr := "user=postgres password=postgres dbname=postgres sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func createTables(db *sql.DB) {
	_, err := db.Exec(`CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		name TEXT NOT NULL,
		email TEXT NOT NULL UNIQUE,
		created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
	);`)
	if err != nil {
		log.Fatalf("Error creating users table: %s", err)
	} else {
		fmt.Println("Users table created successfully")
	}

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS categories (
		id SERIAL PRIMARY KEY,
		name TEXT NOT NULL
	);`)
	if err != nil {
		log.Fatalf("Error creating categories table: %s", err)
	} else {
		fmt.Println("categories table Created successfully")
	}

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS products (
		id SERIAL PRIMARY KEY,
		name TEXT NOT NULL,
		price DECIMAL(10,2),
		category_id INTEGER REFERENCES categories(id),
		created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
	);`)
	if err != nil {
		log.Fatalf("Error creating products table: %s", err)
	} else {
		fmt.Println("Products table Created successfully")
	}

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS orders (
		id SERIAL PRIMARY KEY,
		user_id INTEGER REFERENCES users(id),
		status TEXT NOT NULL,
		created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
	);`)
	if err != nil {
		log.Fatalf("Error creating orders table: %s", err)
	} else {
		fmt.Println("orders table Created successfully")
	}

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS orderDetails (
		id SERIAL PRIMARY KEY,
		order_id INTEGER REFERENCES orders(id),
		product_id INTEGER REFERENCES products(id),
		quantity INTEGER NOT NULL,
		price NUMERIC
	);`)
	if err != nil {
		log.Fatalf("Error creating orderDetails table: %s", err)
	} else {
		fmt.Println("orderDetails table Created successfully")
	}

}
