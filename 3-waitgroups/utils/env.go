package utils

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("FAIL: .env file can not load")
	}
}

func GetEnv(key, fallback string) string {
	val := os.Getenv(key)

	if val == "" {
		return fallback //return fallback value if .env is empty (for secret keys write only "" for not to show it on main project code)
	}
	return val
}
