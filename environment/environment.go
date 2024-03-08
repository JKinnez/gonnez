package environment

import (
	"os"
)

var envKey string

func Init(key string) {
	envKey = key
}

func IsProduction() bool {
	return os.Getenv(envKey) == "production"
}

func IsDevelopment() bool {
	return os.Getenv(envKey) == "development"
}

func IsTest() bool {
	return os.Getenv(envKey) == "test"
}

func IsCI() bool {
	return os.Getenv(envKey) == "ci"
}
