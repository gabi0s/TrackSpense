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
	dbHost := os.Getenv("DB_HOST_USER")
	if dbHost == "" {
		log.Fatal("DB_HOST_USER is not set")
	}

	dbPort := os.Getenv("DB_PORT")
	if dbPort == "" {
		log.Fatal("DB_PORT is not set")
	}

	dbUser := os.Getenv("POSTGRES_USER_USER")
	if dbUser == "" {
		log.Fatal("POSTGRES_USER_USER is not set")
	}

	dbPassword := os.Getenv("POSTGRES_PASSWORD_USER")
	if dbPassword == "" {
		log.Fatal("POSTGRES_PASSWORD_USER is not set")
	}

	dbName := os.Getenv("POSTGRES_DB_USER")
	if dbName == "" {
		log.Fatal("POSTGRES_DB_USER is not set")
	}

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		dbHost, dbUser, dbPassword, dbName, dbPort)

	// effectively initiate the connection with env variables
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{DisableForeignKeyConstraintWhenMigrating: true})
	if err != nil {
		log.Fatal("failed to connect to database:", err)
	}

	log.Println("Successfully connected to the database")
}
