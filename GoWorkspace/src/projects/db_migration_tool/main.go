package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/pseudonative/DevelopmentGo/GoWorkspace/src/projects/db_migration_tool/db"
)

func main() {
	const (
		user     = "postgres"
		password = "mysecretpassword"
		dbname   = "postgres"
	)

	EnsureMigrationDir("migrations")

	db, err := db.ConnectToDB(user, password, dbname)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()
	log.Println("Successfully connnected to database!")
}

func EnsureMigrationDir(dirName string) {
	if _, err := os.Stat(dirName); os.IsNotExist(err) {
		err := os.Mkdir(dirName, 0755)
		if err != nil {
			log.Fatalf("Failed to create migration directory: %v", err)
		}
		log.Printf("Create migration directory: %s", dirName)
	}
}

func CreateMigrationFile(migrationName string) {
	timestamp := time.Now().Format("20060102150405")
	fileName := fmt.Sprintf("%s_%s.sql", timestamp, migrationName)

	content := `-- UP
-- SQL statements for applying the migration go here

-- DOWN
-- SQL statements for rolling back the migration go here
`
	f, err := os.Create("migrations/" + fileName)
	if err != nil {
		log.Fatalf("Failed to create migration file: %v", err)
	}
	defer f.Close()

	_, err = f.WriteString(content)
	if err != nil {
		log.Fatalf("Failed to write to migration file: %v", err)
	}
	log.Printf("Created migration file: %s", fileName)
}
