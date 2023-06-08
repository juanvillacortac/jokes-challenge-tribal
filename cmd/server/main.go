package main

import (
	"fmt"
	"log"
	"os"

	"challenge/pkg/server"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	port := os.Getenv("PORT")
	addr := fmt.Sprintf(":%s", port)

	log.Printf("Listening on %s", addr)
	srv := server.NewServer()

	if err := srv.Listen(addr); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
