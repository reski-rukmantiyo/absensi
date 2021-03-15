/**
 * Taken from https://dev.to/craicoverflow/a-no-nonsense-guide-to-environment-variables-in-go-a2f
 */

package config

import (
	"log"
	"os"
	"os/user"
	"strconv"
	"strings"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}
	if err := godotenv.Load("/app/.env"); err != nil {
		if err := godotenv.Load(usr.HomeDir + "/.env"); err != nil {
			if err := godotenv.Load(".env"); err != nil {
				log.Fatal("No .env file. Application dismissed")
			}
		}
	}
}

// GetEnv : Simple helper function to read an environment or return a default value
func GetEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultVal
}

// GetEnvAsInt : Simple helper function to read an environment variable into integer or return a default value
func GetEnvAsInt(name string, defaultVal int) int {
	valueStr := GetEnv(name, "")
	if value, err := strconv.Atoi(valueStr); err == nil {
		return value
	}

	return defaultVal
}

// GetEnvAsBool : Helper to read an environment variable into a bool or return default value
func GetEnvAsBool(name string, defaultVal bool) bool {
	valStr := GetEnv(name, "")
	if val, err := strconv.ParseBool(valStr); err == nil {
		return val
	}

	return defaultVal
}

// GetEnvAsSlice : Helper to read an environment variable into a string slice or return default value
func GetEnvAsSlice(name string, defaultVal []string, sep string) []string {
	valStr := GetEnv(name, "")

	if valStr == "" {
		return defaultVal
	}

	val := strings.Split(valStr, sep)

	return val
}
