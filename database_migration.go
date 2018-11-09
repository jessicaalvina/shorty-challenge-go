package main

import (
	"github.com/joho/godotenv"
	"log"
)

func init() {

	if godotenv.Load() != nil {
		log.Fatal("Error loading .env file")
	}

}

func main() {

}
