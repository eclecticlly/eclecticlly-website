package env

import (
	"os"

	"github.com/joho/godotenv"
)

var Env map[string]string

func GetEnv(key, def string) string {
	if val, ok := os.LookupEnv(key); ok {
		return val
	}
	return def
}

func SetupEnvFile() {
	env := GetEnv("APP_ENV", "")
	if env == "" {
		err := godotenv.Load()
		if err != nil {
			panic(err)
		}
	}
}
