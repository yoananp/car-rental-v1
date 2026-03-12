package main

import (
	"fmt"
	"log"
	"os"

	database "github.com/yoananp/car-rental-v1/database/sql_migration"
	"github.com/yoananp/car-rental-v1/routes"

	"github.com/joho/godotenv"
)

func main() {
	// Load .env
	if err := godotenv.Load("config/.env"); err != nil {
		fmt.Println("Failed to load .env file")
	} else {
		fmt.Println("Success to load .env file")
	}

	// Build connection string
	psqlInfo := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("PGHOST"),
		os.Getenv("PGPORT"),
		os.Getenv("PGUSER"),
		os.Getenv("PGPASSWORD"),
		os.Getenv("PGDATABASE"),
	)

	// Connect to DB
	DB, err := database.ConnectDB(psqlInfo)
	if err != nil {
		log.Fatal("DB Connection Failed:", err)
	}
	defer DB.Close()

	// Run migrations
	database.DbMigrate(DB)

	// Start HTTP server (bisa pakai Gin via routes)
	router := routes.Route()
	fmt.Println("Server running at http://localhost:8080")
	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
