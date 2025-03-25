package services

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres" // for psql db connection
	"gorm.io/gorm"            // ORM for go
)

// DB is global gorm connection to the database
var DB *gorm.DB

// Intitialize the connection to the database
func InitDB() {
	// gathering environment variables
	dbHostBudget := os.Getenv("DB_HOST_BUDGET")
	if dbHostBudget == "" {
		log.Fatal("DB_HOST_BUDGET is not set")
	}

	dbPortBudget := os.Getenv("DB_PORT")
	if dbPortBudget == "" {
		log.Fatal("DB_PORT is not set")
	}

	dbUserBudget := os.Getenv("POSTGRES_USER_BUDGET")
	if dbUserBudget == "" {
		log.Fatal("POSTGRES_USER_BUDGET is not set")
	}

	dbPasswordBudget := os.Getenv("POSTGRES_PASSWORD_BUDGET")
	if dbPasswordBudget == "" {
		log.Fatal("POSTGRES_PASSWORD_BUDGET is not set")
	}

	dbNameBudget := os.Getenv("POSTGRES_DB_BUDGET")
	if dbNameBudget == "" {
		log.Fatal("POSTGRES_DB_BUDGET is not set")
	}

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		dbHostBudget, dbUserBudget, dbPasswordBudget, dbNameBudget, dbPortBudget)

	// effectively initiate the connection with env. variables
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{DisableForeignKeyConstraintWhenMigrating: true})
	if err != nil {
		log.Fatal("failed to connect to database:", err)
	}

	log.Println("Successfully connected to the database")
}
