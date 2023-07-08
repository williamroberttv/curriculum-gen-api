package config

import (
	"os"

	"github.com/joho/godotenv"
)

// load env file when GO_ENV is not production
func LoadENV() error {
	goEnv := os.Getenv("GO_ENV")
	if goEnv == "" || goEnv == "development" {
		err := godotenv.Load()
		if err != nil {
			return err
		}
	}
	return nil
}