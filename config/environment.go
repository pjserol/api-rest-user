package config

import (
	"log"
)

// Environment environment variable of the application
type Environment struct {
	AppEnvironment string
	IsTestMode     bool
	SuppressLogs   string
	PortNumber     int

	// Database
	TestDatabaseURL string
	AWSRegion       string
	DBEndpoint      string
	DBPort          int
	DBUser          string
	DBName          string
}

// InitEnvironment initialise environment variable of the application
func InitEnvironment() Environment {
	env := Environment{
		AppEnvironment: getEnvString("APP_ENVIRONMENT", "local"), // local, dev, stag, prod.
		IsTestMode:     getEnvBool("IS_TEST_MODE", false),
		SuppressLogs:   getEnvString("SUPPRESS_LOGS", "true"),
		PortNumber:     getEnvInt("PORT_NUMBER", 8080),

		// Database
		TestDatabaseURL: getEnvString("TEST_DATABASE_URL", "DEFAULT-VALUE"),
		AWSRegion:       getEnvString("AWS_REGION", "DEFAULT-VALUE"),
		DBEndpoint:      getEnvString("DATABASE_ENDPOINT", "DEFAULT-VALUE"),
		DBPort:          getEnvInt("DATABASE_PORT", 5432),
		DBUser:          getEnvString("DATABASE_USER", "DEFAULT-VALUE"),
		DBName:          getEnvString("DATABASE_NAME", "DEFAULT-VALUE"),
	}

	if env.AppEnvironment == "" {
		log.Println("error:APP_ENVIRONMENT should not be empty")
	}

	return env
}
