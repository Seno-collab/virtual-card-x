package main

import (
	"log"
	"os"

	"example.com/virtual-card-x/internal/server"
)

func main() {
	dsn := os.Getenv("DATABASE_URL")
	srv, err := server.New(dsn)
	if err != nil {
		log.Fatal(err)
	}
	if err := srv.Run(os.Getenv("ADDR")); err != nil {
		log.Fatal(err)
	}
}
