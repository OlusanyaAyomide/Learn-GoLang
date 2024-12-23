package env

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

// if first variable is not typed , it is auto inferred as the last typed variable
// Return envs
func GetEnvString(key, fallback string) string {
	err := godotenv.Load(".env")

	if err != nil {
		log.Println(err, "here")
		return fallback
	} else {
		log.Println("Env loaded")
	}
	val := os.Getenv(key)

	return val
}

func GetEnvInt(key string, fallback int) int {
	err := godotenv.Load(".env")

	if err != nil {
		return fallback
	} else {
		log.Println("Env loaded")
	}

	val := os.Getenv(key)
	valAsInt, err := strconv.Atoi(val)
	if err != nil {
		return fallback
	}

	return valAsInt
}
