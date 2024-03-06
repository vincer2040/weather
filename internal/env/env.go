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

func GetProductionValue() bool {
	isProd := os.Getenv("PRODUCTION")
	return isProd == "true"
}

func GetDBURL() string {
	if GetProductionValue() {
		return os.Getenv("PROD_DB_URL")
	}
	return os.Getenv("DB_URL")
}

func GetWeatherApiKey() string {
    return os.Getenv("WEATHER_API_KEY")
}
