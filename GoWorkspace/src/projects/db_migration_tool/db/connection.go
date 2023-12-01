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

func RecordMigration(db *sql.DB, migrationName string) error {
	_, err := db.Exec("INSERT INTO applied_migrations (name, applied_at) VALUES ($1, NOW())", migrationName)
	return err
}

func RemoveMigrationRecord(db *sql.DB, migrationName string) error {
	_, err := db.Exec("DELETE FROM applied_migrations WHERE name = $1", migrationName)
	return err
}

func CheckMigrationApplied(db *sql.DB, migrationName string) (bool, error) {
	var count int
	err := db.QueryRow("SELECT COUNT(*) FROM applied_migrations WHERE name = $1", migrationName).Scan(&count)
	return count > 0, err
}
