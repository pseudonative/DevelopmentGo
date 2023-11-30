package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/pseudonative/DevelopmentGo/GoWorkspace/src/projects/db_migration_tool/db"
)

func main() {
	// Define flags for migration file names
	migrationFileName := flag.String("migrate", "", "The name of the migration file to apply")
	rollbackFileName := flag.String("rollback", "", "The name of the migration file to rollback")
	flag.Parse()

	// Database connection configuration
	const (
		user     = "postgres"
		password = "mysecretpassword"
		dbname   = "postgres"
	)

	// Establish database connection
	dbConn, err := db.ConnectToDB(user, password, dbname)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer dbConn.Close()
	log.Println("Successfully connected to the database!")

	// Apply migration
	if *migrationFileName != "" {
		upCommands, _, err := readMigrationFiles(*migrationFileName)
		if err != nil {
			log.Fatalf("Error reading migration file: %v", err)
		}
		err = db.ApplyMigration(dbConn, upCommands)
		if err != nil {
			log.Fatalf("Error applying migration: %v", err)
		}
		log.Println("Migration applied successfully")
	}

	// Rollback migration
	if *rollbackFileName != "" {
		_, downCommands, err := readMigrationFiles(*rollbackFileName)
		if err != nil {
			log.Fatalf("Error reading migration file for rollback: %v", err)
		}
		err = db.RollbackMigration(dbConn, downCommands)
		if err != nil {
			log.Fatalf("Error rolling back migration: %v", err)
		}
		log.Println("Migration rolled back successfully")
	}

	// Example for creating a migration file
	EnsureMigrationDir("migrations")
	CreateMigrationFile("test_migration")
	log.Println("Migration file creation process completed.")
}

// EnsureMigrationDir checks and creates the migrations directory if it does not exist
func EnsureMigrationDir(dirName string) {
	if _, err := os.Stat(dirName); os.IsNotExist(err) {
		err := os.Mkdir(dirName, 0755)
		if err != nil {
			log.Fatalf("Failed to create migration directory: %v", err)
		}
		log.Printf("Create migration directory: %s", dirName)
	}
}

// CreateMigrationFile creates a new migration file with a specified name
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

// readMigrationFiles reads a migration file and separates it into up and down sections
func readMigrationFiles(fileName string) (string, string, error) {
	file, err := os.Open("migrations/" + fileName)
	if err != nil {
		return "", "", err
	}
	defer file.Close()

	var upCommands, downCommands []string
	var isDownSection bool

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "-- UP") {
			isDownSection = false
			continue
		} else if strings.Contains(line, "-- DOWN") {
			isDownSection = true
			continue
		}
		if isDownSection {
			downCommands = append(downCommands, line)
		} else {
			upCommands = append(upCommands, line)
		}
	}

	if err := scanner.Err(); err != nil {
		return "", "", err
	}

	return strings.Join(upCommands, "\n"), strings.Join(downCommands, "\n"), nil
}
