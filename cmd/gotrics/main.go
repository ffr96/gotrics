package main

import (
	"flag"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gotrics.cl/internal/api"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	addr := flag.String("addr", "localhost:4000", "HTTP network address")
	dsn := flag.String("dsn", os.Getenv("DATABASE_URL"), "MySQL data source name")

	flag.Parse()

	api.Main(addr, dsn)
}
