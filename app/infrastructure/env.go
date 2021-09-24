package infrastructure

import (
	"log"

	"github.com/directoryxx/fiber-clean-template/app/usecases"
	"github.com/joho/godotenv"
)

// Load is load configs from a env file.
func Load(logger usecases.Logger) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
