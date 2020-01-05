package yetconfig

import (
	"log"

	"github.com/JeremyLoy/config"
)

// LoadEnvConfig loads the configuration from .env files and/or from OS environment.
// The current directory can be respresented by ".".
//
// Example:
//  type Config struct {
//    Port int
//  }
//
//  cfg := Config{}
//  LoadEnvConfig(&cfg, ".")
func LoadEnvConfig(targetConfig interface{}, pathToEnvFiles string) {
	envFile := selectEnvFileByEnvironment()

	envFilePath := mergeEnvFilePath(envFile, pathToEnvFiles)
	customEnvFilePath := mergeEnvFilePath(envFileCustom, pathToEnvFiles)

	if err := config.From(envFilePath).From(customEnvFilePath).FromEnv().To(targetConfig); err != nil {
		log.Printf("WARNING: some errors occured on config load: %s", err.Error())
	}
}
