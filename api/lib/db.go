package lib

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func GetDBConnection() *gorm.DB {
	var db *gorm.DB
	var err error

	env, _ := LoadEnvVariables(".env")

	dsn := env["DB_USERNAME"] + ":" + env["DB_PASSWORD"] + "@tcp(" + env["DB_URL"] + ":" + env["DB_PORT"] + ")/" + env["DB_NAME"] + "?charset=utf8mb4&parseTime=True&loc=Local"

	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	return db
}

func LoadEnvVariables(filePath string) (map[string]string, error) {
	err := godotenv.Load(filePath)
	if err != nil {
		return nil, fmt.Errorf("error loading .env file: %w", err)
	}

	envVars := make(map[string]string)

	for _, env := range os.Environ() {
		key, value, err := parseEnv(env)

		if err != nil {
			return nil, fmt.Errorf("error parsing environment variable: %w", err)
		}

		envVars[key] = value
	}

	return envVars, nil
}

func parseEnv(env string) (key, value string, err error) {
	pair := strings.SplitN(env, "=", 2)

	if len(pair) != 2 {
		return "", "", fmt.Errorf("malformed environment variable: %s", env)
	}

	return pair[0], pair[1], nil
}
