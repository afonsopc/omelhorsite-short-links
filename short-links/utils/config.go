package utils

import (
	"log"
	"os"
)

type ApiConfiguration struct {
	Endpoint string
}

type AccountsServiceConfiguration struct {
	Endpoint string
}

type DatabaseConfiguration struct {
	Uri string
}

func getEnvironmentVariable(value string) string {
	variable := os.Getenv(value)

	if variable == "" {
		log.Fatalf("%s environment variable not found. Define it 😠", value)
	}

	return variable
}

func GetApiConfiguration() ApiConfiguration {
	return ApiConfiguration{
		Endpoint: getEnvironmentVariable("ENDPOINT"),
	}
}

func GetAccountsServiceConfiguration() AccountsServiceConfiguration {
	return AccountsServiceConfiguration{
		Endpoint: getEnvironmentVariable("ACCOUNTS_SERVICE_ENDPOINT"),
	}
}

func GetDatabaseConfiguration() DatabaseConfiguration {
	return DatabaseConfiguration{
		Uri: getEnvironmentVariable("DATABASE_URI"),
	}
}

func CheckAllConfigurationVariables() {
	GetApiConfiguration()
	GetAccountsServiceConfiguration()
	GetDatabaseConfiguration()
}
