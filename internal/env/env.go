package env

import (
	"os"

	"github.com/joho/godotenv"
)

func EnvInit() error {
    err := godotenv.Load()
    return err
}

func GetSessionSecret() string {
    return os.Getenv("SESSION_SECRET")
}
