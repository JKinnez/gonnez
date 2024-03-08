package tokenizer

import (
	"fmt"
	"os"
	"time"
	_ "time/tzdata"

	"github.com/go-playground/validator/v10"
)

const issuer = "gonnez-tokenizer-issuer"
const defaultSymetricKeyEnvName = "GONNEZ_SYMMETRIC_KEY"
const defaultPrivateKeyEnvName = "GONNEZ_PRIVATE_KEY"
const defaultPublicKeyEnvName = "GONNEZ_PUBLIC_KEY"
const bearerPrefix = "Bearer "
const emptyString = ""

var validate *validator.Validate

func buildBearer(token string) string {
	return fmt.Sprintf("%s%s", bearerPrefix, token)
}

func splitBearer(bearer string) string {
	return bearer[len(bearerPrefix):]
}

func setIssuer(value string) string {
	if value == emptyString {
		return issuer
	}

	return value
}

func validation(record any) (err error) {
	validate = validator.New()
	err = validate.Struct(record)
	return
}

func symetrickey(value string) string {
	if value == emptyString {
		return os.Getenv(defaultSymetricKeyEnvName)
	}

	return os.Getenv(value)
}

func privatekey(value string) string {
	if value == emptyString {
		return os.Getenv(defaultPrivateKeyEnvName)
	}

	return os.Getenv(value)
}

func publickey(value string) string {
	if value == emptyString {
		return os.Getenv(defaultPublicKeyEnvName)
	}

	return os.Getenv(value)
}

func now() time.Time {
	return time.Now().UTC()
}

func toLocale(current time.Time, location string) (currentLocale time.Time, err error) {
	if location == emptyString {
		currentLocale = current
		return
	}

	loc, err := time.LoadLocation(location)
	if err != nil {
		return
	}

	currentLocale = current.In(loc)
	return
}

func expiracy(current time.Time, expiration int) time.Time {
	return current.Add(time.Duration(expiration) * time.Hour)
}
