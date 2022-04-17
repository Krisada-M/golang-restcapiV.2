package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func EnvMongoURI() string {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	return os.Getenv("MONGOURI")
}

func EnvDBname() string {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	return os.Getenv("DB_NAME")
}

func EnvCollectionname() string {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	return os.Getenv("COLLECTION_NAME")
}
