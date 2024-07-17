package lib

import (
	"os"
	"strconv"
	"strings"

	"github.com/joho/godotenv"
)

func GetEnvInt(value string) int {
	str := os.Getenv(value)
	intValue, _ := strconv.Atoi(str)
	return intValue
}

func GetEnvString(value string) string {
	return os.Getenv(value)
}

func GetEnvBool(value string) bool {
	str := os.Getenv(value)
	return strings.ToUpper(str) == "TRUE"
}

func LoadEnv() error {
	err := godotenv.Load()
	if err != nil {
		panic(err.Error())
	}
	return nil
}
