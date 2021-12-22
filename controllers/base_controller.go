package controllers

import (
	"log"

	"github.com/joho/godotenv"
)

func init() {
	// load .env file
	err := godotenv.Load()

	if err != nil {
		log.Println("[!] load dot.env error ->", err.Error())
		return
	}

	// store .env variable:

	// Load Login Provider Configuration
	LoadLoginProviderConfiguration()

	// Load ORY Kratos Configuration
	LoadKratosConfiguration()

	// Load ORY Hydra Configuration
	LoadHydraConfiguration()
}
