package yetenv

import (
	"os"
	"strings"
)

const (
	environmentVariableName            = "ENVIRONMENT"
	environmentVariableValueProduction = "production"
	environmentVariableValueStaging    = "staging"
)

type Environment string

const (
	Production Environment = "production"
	Staging    Environment = "staging"
	Develop    Environment = "develop"
)

func GetEnvironment() Environment {
	envRaw := os.Getenv(environmentVariableName)
	env := strings.ToLower(envRaw)

	switch env {
	case environmentVariableValueProduction:
		return Production
	case environmentVariableValueStaging:
		return Staging
	}

	return Develop
}
