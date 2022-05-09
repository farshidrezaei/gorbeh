package libraryDotEnv

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

func LoadEnvironmentVariables() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Some error occured on loading envirement variables. Err: %s", err)
	}
}

func Env(variable string, defaultValue string) string {

	if value := os.Getenv(variable); value != "" {
		return value
	}
	return defaultValue
}
