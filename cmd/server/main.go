package main

import (
	"fmt"
	"log"
	"os"

	"challenge/pkg/server"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	addr := fmt.Sprintf(":%s", port)

	log.Printf("Listening on %s", addr)
	srv := server.NewServer()

	if err := srv.Listen(addr); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
