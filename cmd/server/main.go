package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/King0625/grpc-task-app/internal/db"
	"github.com/joho/godotenv"
)

func main() {
	env := os.Getenv("ENV")
	if env != "production" {
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
		}
	}
	dsn := os.Getenv("POSTGRES_DSN")

	if err := db.RunMigration(dsn); err != nil {
		log.Fatalf("run migration error: %v", err)
	}

	conn, err := db.InitPostgres(dsn)
	if err != nil {
		log.Fatalf("Cannot init postgres conn: %v", err)
	}
	defer conn.Close(context.Background())
	fmt.Println("good to go!")
}
