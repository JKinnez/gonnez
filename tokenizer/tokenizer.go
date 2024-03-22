package tokenizer

import (
	"fmt"
	"os"
	"time"
	_ "time/tzdata"

	"github.com/go-playground/validator/v10"
)

const (
	issuer                    = "gonnez-tokenizer-issuer"
	defaultSymetricKeyEnvName = "GONNEZ_SYMMETRIC_KEY"
	defaultPrivateKeyEnvName  = "GONNEZ_PRIVATE_KEY"
	defaultPublicKeyEnvName   = "GONNEZ_PUBLIC_KEY"
	bearerPrefix              = "Bearer "
	emptyString               = ""
	keyLenght                 = 32
)

type Duration int64

var (
	validate *validator.Validate
	hours    Duration = 24
)

var Durations = struct {
	OneDay      Duration
	TwoDays     Duration
	ThreeDays   Duration
	OneWeek     Duration
	TwoWeeks    Duration
	ThreeWeeks  Duration
	OneMonth    Duration
	TwoMonths   Duration
	ThreeMonths Duration
}{
	OneDay:      hours,
	TwoDays:     2 * hours,
	ThreeDays:   3 * hours,
	OneWeek:     7 * hours,
	TwoWeeks:    14 * hours,
	ThreeWeeks:  21 * hours,
	OneMonth:    30 * hours,
	TwoMonths:   60 * hours,
	ThreeMonths: 90 * hours,
}

type Config struct {
	Audience           string
	Footer             *string
	Issuer             string
	Location           string
	PublicKeyEnvName   string
	PrivateKeyEnvName  string
	SymetricKeyEnvName string
	// revive:disable:struct-tag
	expiration Duration `validate:"min=1"`
	subject    string   `validate:"required"`
	token      string   `validate:"required"`
	// revive:enable:struct-tag
}

type Tokenizer struct {
	Config
}

func New(options Config) *Tokenizer {
	return &Tokenizer{
		options,
	}
}

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

func expiracy(current time.Time, expiration Duration) time.Time {
	return current.Add(time.Duration(expiration) * time.Hour)
}
