package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {

	fmt.Println("Successfully connected to MongoDB!")

	godotenv.Load()

	portString := os.Getenv("PORT")

	if portString == "" {
		log.Fatal("PORT environment variable is not set")
	}

	fmt.Printf("Server is running on port %s\n", portString)

}
