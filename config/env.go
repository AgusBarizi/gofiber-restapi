package config

import (
	"github.com/joho/godotenv"
	"m3gaplazma/gofiber-restapi/exception"
	"os"
)

func Env(key string) string {
	err := godotenv.Load(".env")
	exception.PanicIfError(err)
	return os.Getenv(key)
}
