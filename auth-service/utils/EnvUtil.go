package utils

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func GetEnvVar(key string) string {
	if os.Getenv("ENVIRONMENT") != "production" {
		err := godotenv.Load()
		if err != nil {
			fmt.Println("Error loading.env file", err)
		}
	}
	return os.Getenv(key)
}
