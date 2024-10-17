package environment

var env string

func Init(key string) {
	env = key
}

func IsProduction() bool {
	return env == "production"
}

func IsDevelopment() bool {
	return env == "development"
}

func IsTest() bool {
	return env == "test"
}

func IsCI() bool {
	return env == "ci"
}
