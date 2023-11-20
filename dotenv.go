package dotenv

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"github.com/pkg/errors"
)

// Load auto-load env variables from .env file in the root of project
func Load() error {
	err := godotenv.Load()
	if err != nil {
		return errors.Wrap(err, "Error loading .env file")
	}

	return nil
}

func GetString(varName string, defaultValue string) string {
	if envVal, ok := os.LookupEnv(varName); ok {
		return envVal
	}

	return defaultValue
}

func GetInt(varName string, defaultValue int) int {
	envVal, ok := os.LookupEnv(varName)

	if !ok {
		return defaultValue
	}

	if num, err := strconv.Atoi(envVal); err == nil {
		return num
	}

	return defaultValue
}
