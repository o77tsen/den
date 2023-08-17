package main

import (
	"log"

	"github.com/o77tsen/den/initializers"
	"github.com/o77tsen/den/models"
)

func init() {
	initializers.LoadEnv()
	initializers.Connect()
}

func main() {
	err := initializers.DB.AutoMigrate(&models.Message{})
	if err != nil {
		log.Fatalf("Error during db migration: %v", err)
	}
}