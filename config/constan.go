package config

import (
	"os"

	"github.com/joho/godotenv"
)

func Path_url() string {
	err := godotenv.Load()
	if err != nil {
		panic("Failed to Load env File")
	}
	path := os.Getenv("PATH_API_BACKEND")

	return path
}
