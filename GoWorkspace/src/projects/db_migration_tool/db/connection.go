package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
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
		return nil, err
	}
	log.Println("Successfully connected to the databse")
	return db, nil
}

func ApplyMigration(db *sql.DB, upCommands string) error {
	_, err := db.Exec(upCommands)
	if err != nil {
		return err
	}
	return nil
}

func RollbackMigration(db *sql.DB, downCommands string) error {
	_, err := db.Exec(downCommands)
	if err != nil {
		return err
	}
	return nil
}
