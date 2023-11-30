package db

import (
	"database/sql"
	"fmt"
	"log"
)

func ConnectToDB(user, password, dbname string) (*sql.DB, error) {
	connStr := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", user, password, dbname)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Printf("Error opening database: %v", err)
	}
	err = db.Ping()
	if err != nil {
		log.Printf("Error connecting to database: %v", err)
	}
	log.Println("Successfully connected to the databsed")
	return db, nil
}
